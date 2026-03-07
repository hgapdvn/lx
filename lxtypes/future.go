package lxtypes

import (
	"context"
	"errors"
	"sync"
)

// Future represents a value that will be available in the future.
// It supports type transformation through the Then function for sequential
// operations, and can be combined with other futures for parallel execution.
//
// Example (sequential):
//
//	cardFuture := FutureDo(func() (int, error) {
//	    return getUserId(), nil
//	})
//	cardFuture = Then(cardFuture, func(userId int) (User, error) {
//	    return fetchUser(userId)
//	})
//	cardFuture = Then(cardFuture, func(user User) (Card, error) {
//	    return fetchCard(user.CardId)
//	})
//	card, err := cardFuture.Get(ctx)
//
// Example (parallel):
//
//	s1 := FutureDo(func() (Data, error) { return fetchService1() })
//	s2 := FutureDo(func() (Data, error) { return fetchService2() })
//	s3 := FutureDo(func() (Data, error) { return fetchService3() })
//	allData := FutureAll(s1, s2, s3)
//	results, err := allData.Get(ctx)
type Future[T any] interface {
	// Get blocks until the computation completes and returns the result.
	//
	// If the provided context is canceled or reaches its deadline before
	// the computation completes, Get returns the context error.
	Get(ctx context.Context) (T, error)
}

// Note: Then is provided as a standalone function (not a method) because
// Go interfaces cannot have methods with type parameters.
// Use: Then(future, transformFn) instead of future.Then(transformFn)

// future is the internal implementation of Future[T].
// Lock-free design using only channel synchronization.
type future[T any] struct {
	fn    func() (T, error) // Function to execute
	value T                 // Result value
	err   error             // Result error
	done  chan struct{}     // Closed when computation completes
	once  sync.Once         // Ensures fn runs only once
}

// Get blocks until the computation completes and returns the result.
// Respects context cancellation and deadlines.
//
// If the context is cancelled or times out before the future completes,
// returns the zero value of T and the context error.
//
// Note: The future's computation continues running in the background even
// if the context is cancelled. The cancellation only affects this Get call.
func (f *future[T]) Get(ctx context.Context) (T, error) {
	select {
	case <-f.done:
		// Future completed (successfully or with error)
		return f.value, f.err
	case <-ctx.Done():
		// Context cancelled or deadline exceeded
		var zero T
		err := ctx.Err()
		if err == nil {
			// This should never happen according to Go's context contract,
			// but defend against it anyway
			err = context.Canceled
		}
		return zero, err
	}
}

// exec executes the function and stores the result.
// Lock-free: writes happen-before channel close.
func (f *future[T]) exec() {
	f.once.Do(func() {
		f.value, f.err = f.fn()
		close(f.done) // Signals completion (happens-after guarantee)
	})
}

// FutureThen Then creates a new Future that runs after the current Future successfully
// completes, transforming the result from type T to type U.
//
// This is a standalone function (not a method) because Go interfaces cannot
// have methods with type parameters.
//
// If the parent Future completes with an error, the error is propagated and
// the transformation function is not executed.
//
// Context cancellation is handled by Get() - if you call Get(ctx) on the
// chained future with a cancelled context, it will return immediately with
// the context error, even if the parent is still running.
//
// Example:
//
//	// Transform int -> User -> Card
//	userIdFuture := FutureDo(func() (int, error) {
//	    return getUserId(), nil
//	})
//	userFuture := FutureThen(userIdFuture, func(userId int) (User, error) {
//	    return fetchUser(userId)
//	})
//	cardFuture := FutureThen(userFuture, func(user User) (Card, error) {
//	    return fetchCard(user.CardId)
//	})
//	card, err := cardFuture.Get(ctx)
func FutureThen[T, U any](parent Future[T], fn func(T) (U, error)) Future[U] {
	// Cast to access internal done channel for efficient waiting
	// This avoids blocking on context.Background() in Get()
	parentImpl := parent.(*future[T])

	next := &future[U]{
		fn: func() (U, error) {
			// Wait for parent's done channel directly
			// This makes the child future's Get() context-cancellable
			<-parentImpl.done

			// If parent failed, propagate error
			if parentImpl.err != nil {
				var zero U
				return zero, parentImpl.err
			}

			// Transform T -> U
			return fn(parentImpl.value)
		},
		done: make(chan struct{}),
	}

	// Start the chained future immediately
	go next.exec()
	return next
}

// FutureDo creates a Future that executes the given function asynchronously.
// The computation starts immediately in a background goroutine (hot start).
//
// Example:
//
//	future := FutureDo(func() (string, error) {
//	    resp, err := http.Get("https://api.example.com")
//	    if err != nil {
//	        return "", err
//	    }
//	    defer resp.Body.Close()
//	    data, _ := io.ReadAll(resp.Body)
//	    return string(data), nil
//	})
//	result, err := future.Get(ctx)
func FutureDo[T any](fn func() (T, error)) Future[T] {
	f := &future[T]{
		fn:   fn,
		done: make(chan struct{}),
	}
	go f.exec()
	return f
}

// FutureOf creates a Future that is already completed with the given value.
// No goroutine is started - Get() returns immediately.
//
// Example:
//
//	future := FutureOf(42)
//	value, _ := future.Get(ctx) // Returns 42 immediately
func FutureOf[T any](value T) Future[T] {
	f := &future[T]{
		value: value,
		done:  make(chan struct{}),
	}
	close(f.done)
	return f
}

// FutureError creates a Future that is already completed with an error.
// No goroutine is started - Get() returns the error immediately.
//
// Example:
//
//	future := FutureError[int](errors.New("failed"))
//	_, err := future.Get(ctx) // Returns error immediately
func FutureError[T any](err error) Future[T] {
	f := &future[T]{
		err:  err,
		done: make(chan struct{}),
	}
	close(f.done)
	return f
}

// FutureAll executes multiple futures of the same type concurrently and
// returns a future containing all results as a slice.
// If any future fails, returns the first error encountered.
//
// All futures are executed in parallel regardless of errors, but only
// the first error is returned.
//
// The returned future respects context cancellation - if you call Get(ctx)
// with a cancelled or timed-out context, it returns immediately without
// waiting for all futures to complete.
//
// Example:
//
//	service1 := FutureDo(func() (Data, error) { return fetchService1() })
//	service2 := FutureDo(func() (Data, error) { return fetchService2() })
//	service3 := FutureDo(func() (Data, error) { return fetchService3() })
//
//	allData := FutureAll(service1, service2, service3)
//	results, err := allData.Get(ctx) // []Data{data1, data2, data3}
//
//	// Transform combined results
//	//response := Then(allData, func(data []Data) (Response, error) {
//	//    return combineData(data), nil
//	//})
func FutureAll[T any](futures ...Future[T]) Future[[]T] {
	// Cast all futures to access their done channels
	futureImpls := make([]*future[T], len(futures))
	for i, f := range futures {
		futureImpls[i] = f.(*future[T])
	}

	return FutureDo(func() ([]T, error) {
		results := make([]T, len(futures))
		errs := make([]error, len(futures))
		var wg sync.WaitGroup

		for i, f := range futureImpls {
			wg.Add(1)
			go func(index int, future *future[T]) {
				defer wg.Done()
				// Wait on done channel directly (context-cancellable via outer Get)
				<-future.done
				results[index] = future.value
				errs[index] = future.err
			}(i, f)
		}

		wg.Wait()

		// Return first error encountered
		for _, err := range errs {
			if err != nil {
				return nil, err
			}
		}

		return results, nil
	})
}

type anyResult[T any] struct {
	idx   int
	value T
	err   error
}

// FutureAny returns a Future that completes with the first successfully
// completed child's value (the first child whose error is nil). If all
// provided futures fail, FutureAny returns the first encountered error
// according to the input order of futures.
//
// If no futures are provided, FutureAny returns a failed future immediately.
//
// The returned future respects context cancellation when Get(ctx) is called.
// Child futures are not cancelled; they continue running in the background.
func FutureAny[T any](futures ...Future[T]) Future[T] {
	if len(futures) == 0 {
		return FutureError[T](errors.New("lxtypes: no futures provided"))
	}

	futureImpls := make([]*future[T], len(futures))
	for i, f := range futures {
		futureImpls[i] = f.(*future[T])
	}

	return FutureDo(func() (T, error) {
		// Buffered channel to avoid blocking sends if we return early
		ch := make(chan anyResult[T], len(futureImpls))

		for i, fut := range futureImpls {
			// capture index for deterministic error ordering
			go func(index int, ff *future[T]) {
				<-ff.done
				ch <- anyResult[T]{idx: index, value: ff.value, err: ff.err}
			}(i, fut)
		}

		// Store errors by input index so that if all fail we return the first
		// error according to the original input order.
		errs := make([]error, len(futureImpls))

		for i := 0; i < len(futureImpls); i++ {
			r := <-ch
			if r.err == nil {
				return r.value, nil
			}
			errs[r.idx] = r.err
		}

		// All failed - return first non-nil error by input order
		for _, e := range errs {
			if e != nil {
				var zero T
				return zero, e
			}
		}

		// Should not happen, but return generic error
		var zero T
		return zero, errors.New("lxtypes: unknown error in FutureAny")
	})
}

// FutureJoin2 executes two futures concurrently and combines their results
// into a Pair. Returns an error if either future fails.
//
// The returned future respects context cancellation when Get(ctx) is called.
//
// Example:
//
//	user := FutureDo(func() (User, error) { return fetchUser() })
//	config := FutureDo(func() (Config, error) { return fetchConfig() })
//
//	combined := FutureJoin2(user, config)
//	result, err := combined.Get(ctx)
//	// result.First = User, result.Second = Config
//
//	// Transform combined result
//	response := Then(combined, func(pair Pair[User, Config]) (Response, error) {
//	    return buildResponse(pair.First, pair.Second), nil
//	})
func FutureJoin2[T, U any](f1 Future[T], f2 Future[U]) Future[Pair[T, U]] {
	// Cast to access done channels
	f1Impl := f1.(*future[T])
	f2Impl := f2.(*future[U])

	return FutureDo(func() (Pair[T, U], error) {
		var wg sync.WaitGroup

		wg.Add(2)
		go func() {
			defer wg.Done()
			<-f1Impl.done
		}()
		go func() {
			defer wg.Done()
			<-f2Impl.done
		}()
		wg.Wait()

		if f1Impl.err != nil {
			return Pair[T, U]{}, f1Impl.err
		}
		if f2Impl.err != nil {
			return Pair[T, U]{}, f2Impl.err
		}

		return NewPair(f1Impl.value, f2Impl.value), nil
	})
}

// FutureJoin3 executes three futures concurrently and combines their results
// into a Triple. Returns an error if any future fails.
//
// The returned future respects context cancellation when Get(ctx) is called.
//
// Example:
//
//	userFuture := FutureDo(func() (User, error) { return fetchUser() })
//	configFuture := FutureDo(func() (Config, error) { return fetchConfig() })
//	statsFuture := FutureDo(func() (Stats, error) { return fetchStats() })
//
//	combined := FutureJoin3(userFuture, configFuture, statsFuture)
//	result, err := combined.Get(ctx)
//	// result.First = User, result.Second = Config, result.Third = Stats
func FutureJoin3[T, U, V any](f1 Future[T], f2 Future[U], f3 Future[V]) Future[Triple[T, U, V]] {
	// Cast to access done channels
	f1Impl := f1.(*future[T])
	f2Impl := f2.(*future[U])
	f3Impl := f3.(*future[V])

	return FutureDo(func() (Triple[T, U, V], error) {
		var wg sync.WaitGroup

		wg.Add(3)
		go func() {
			defer wg.Done()
			<-f1Impl.done
		}()
		go func() {
			defer wg.Done()
			<-f2Impl.done
		}()
		go func() {
			defer wg.Done()
			<-f3Impl.done
		}()
		wg.Wait()

		if f1Impl.err != nil {
			return Triple[T, U, V]{}, f1Impl.err
		}
		if f2Impl.err != nil {
			return Triple[T, U, V]{}, f2Impl.err
		}
		if f3Impl.err != nil {
			return Triple[T, U, V]{}, f3Impl.err
		}

		return NewTriple(f1Impl.value, f2Impl.value, f3Impl.value), nil
	})
}

// FutureJoin4 executes four futures concurrently and combines their results
// into a Quad. Returns an error if any future fails.
//
// The returned future respects context cancellation when Get(ctx) is called.
func FutureJoin4[T, U, V, W any](f1 Future[T], f2 Future[U], f3 Future[V], f4 Future[W]) Future[Quad[T, U, V, W]] {
	// Cast to access done channels
	f1Impl := f1.(*future[T])
	f2Impl := f2.(*future[U])
	f3Impl := f3.(*future[V])
	f4Impl := f4.(*future[W])

	return FutureDo(func() (Quad[T, U, V, W], error) {
		var wg sync.WaitGroup

		wg.Add(4)
		go func() {
			defer wg.Done()
			<-f1Impl.done
		}()
		go func() {
			defer wg.Done()
			<-f2Impl.done
		}()
		go func() {
			defer wg.Done()
			<-f3Impl.done
		}()
		go func() {
			defer wg.Done()
			<-f4Impl.done
		}()
		wg.Wait()

		if f1Impl.err != nil {
			return Quad[T, U, V, W]{}, f1Impl.err
		}
		if f2Impl.err != nil {
			return Quad[T, U, V, W]{}, f2Impl.err
		}
		if f3Impl.err != nil {
			return Quad[T, U, V, W]{}, f3Impl.err
		}
		if f4Impl.err != nil {
			return Quad[T, U, V, W]{}, f4Impl.err
		}

		return NewQuad(f1Impl.value, f2Impl.value, f3Impl.value, f4Impl.value), nil
	})
}

// FutureJoin5 executes five futures concurrently and combines their results
// into a Tuple5. Returns an error if any future fails.
//
// The returned future respects context cancellation when Get(ctx) is called.
//
// Example:
//
//	user := FutureDo(func() (User, error) { return fetchUser() })
//	orders := FutureDo(func() ([]Order, error) { return fetchOrders() })
//	payment := FutureDo(func() (Payment, error) { return fetchPayment() })
//	inventory := FutureDo(func() (Inventory, error) { return fetchInventory() })
//	recommendations := FutureDo(func() ([]Product, error) { return fetchRecommendations() })
//
//	combined := FutureJoin5(user, orders, payment, inventory, recommendations)
//	result, err := combined.Get(ctx)
//	// Access: result.V1, result.V2, result.V3, result.V4, result.V5
func FutureJoin5[T1, T2, T3, T4, T5 any](f1 Future[T1], f2 Future[T2], f3 Future[T3], f4 Future[T4], f5 Future[T5]) Future[Tuple5[T1, T2, T3, T4, T5]] {
	impl1 := f1.(*future[T1])
	impl2 := f2.(*future[T2])
	impl3 := f3.(*future[T3])
	impl4 := f4.(*future[T4])
	impl5 := f5.(*future[T5])

	return FutureDo(func() (Tuple5[T1, T2, T3, T4, T5], error) {
		var wg sync.WaitGroup
		wg.Add(5)

		go func() { defer wg.Done(); <-impl1.done }()
		go func() { defer wg.Done(); <-impl2.done }()
		go func() { defer wg.Done(); <-impl3.done }()
		go func() { defer wg.Done(); <-impl4.done }()
		go func() { defer wg.Done(); <-impl5.done }()

		wg.Wait()

		if impl1.err != nil {
			return Tuple5[T1, T2, T3, T4, T5]{}, impl1.err
		}
		if impl2.err != nil {
			return Tuple5[T1, T2, T3, T4, T5]{}, impl2.err
		}
		if impl3.err != nil {
			return Tuple5[T1, T2, T3, T4, T5]{}, impl3.err
		}
		if impl4.err != nil {
			return Tuple5[T1, T2, T3, T4, T5]{}, impl4.err
		}
		if impl5.err != nil {
			return Tuple5[T1, T2, T3, T4, T5]{}, impl5.err
		}

		return NewTuple5(impl1.value, impl2.value, impl3.value, impl4.value, impl5.value), nil
	})
}

// FutureJoin6 executes six futures concurrently and combines their results
// into a Tuple6. Returns an error if any future fails.
//
// The returned future respects context cancellation when Get(ctx) is called.
//
// Example:
//
//	f1 := FutureDo(func() (int, error) { return 1, nil })
//	f2 := FutureDo(func() (string, error) { return "two", nil })
//	f3 := FutureDo(func() (bool, error) { return true, nil })
//	f4 := FutureDo(func() (float64, error) { return 4.0, nil })
//	f5 := FutureDo(func() ([]int, error) { return []int{5}, nil })
//	f6 := FutureDo(func() (rune, error) { return 'a', nil })
//
//	combined := FutureJoin6(f1, f2, f3, f4, f5, f6)
//	result, err := combined.Get(ctx)
func FutureJoin6[T1, T2, T3, T4, T5, T6 any](f1 Future[T1], f2 Future[T2], f3 Future[T3], f4 Future[T4], f5 Future[T5], f6 Future[T6]) Future[Tuple6[T1, T2, T3, T4, T5, T6]] {
	impl1 := f1.(*future[T1])
	impl2 := f2.(*future[T2])
	impl3 := f3.(*future[T3])
	impl4 := f4.(*future[T4])
	impl5 := f5.(*future[T5])
	impl6 := f6.(*future[T6])

	return FutureDo(func() (Tuple6[T1, T2, T3, T4, T5, T6], error) {
		var wg sync.WaitGroup
		wg.Add(6)

		go func() { defer wg.Done(); <-impl1.done }()
		go func() { defer wg.Done(); <-impl2.done }()
		go func() { defer wg.Done(); <-impl3.done }()
		go func() { defer wg.Done(); <-impl4.done }()
		go func() { defer wg.Done(); <-impl5.done }()
		go func() { defer wg.Done(); <-impl6.done }()

		wg.Wait()

		if impl1.err != nil {
			return Tuple6[T1, T2, T3, T4, T5, T6]{}, impl1.err
		}
		if impl2.err != nil {
			return Tuple6[T1, T2, T3, T4, T5, T6]{}, impl2.err
		}
		if impl3.err != nil {
			return Tuple6[T1, T2, T3, T4, T5, T6]{}, impl3.err
		}
		if impl4.err != nil {
			return Tuple6[T1, T2, T3, T4, T5, T6]{}, impl4.err
		}
		if impl5.err != nil {
			return Tuple6[T1, T2, T3, T4, T5, T6]{}, impl5.err
		}
		if impl6.err != nil {
			return Tuple6[T1, T2, T3, T4, T5, T6]{}, impl6.err
		}

		return NewTuple6(impl1.value, impl2.value, impl3.value, impl4.value, impl5.value, impl6.value), nil
	})
}

// FutureJoin7 executes seven futures concurrently and combines their results
// into a Tuple7. Returns an error if any future fails.
//
// The returned future respects context cancellation when Get(ctx) is called.
func FutureJoin7[T1, T2, T3, T4, T5, T6, T7 any](f1 Future[T1], f2 Future[T2], f3 Future[T3], f4 Future[T4], f5 Future[T5], f6 Future[T6], f7 Future[T7]) Future[Tuple7[T1, T2, T3, T4, T5, T6, T7]] {
	impl1 := f1.(*future[T1])
	impl2 := f2.(*future[T2])
	impl3 := f3.(*future[T3])
	impl4 := f4.(*future[T4])
	impl5 := f5.(*future[T5])
	impl6 := f6.(*future[T6])
	impl7 := f7.(*future[T7])

	return FutureDo(func() (Tuple7[T1, T2, T3, T4, T5, T6, T7], error) {
		var wg sync.WaitGroup
		wg.Add(7)

		go func() { defer wg.Done(); <-impl1.done }()
		go func() { defer wg.Done(); <-impl2.done }()
		go func() { defer wg.Done(); <-impl3.done }()
		go func() { defer wg.Done(); <-impl4.done }()
		go func() { defer wg.Done(); <-impl5.done }()
		go func() { defer wg.Done(); <-impl6.done }()
		go func() { defer wg.Done(); <-impl7.done }()

		wg.Wait()

		if impl1.err != nil {
			return Tuple7[T1, T2, T3, T4, T5, T6, T7]{}, impl1.err
		}
		if impl2.err != nil {
			return Tuple7[T1, T2, T3, T4, T5, T6, T7]{}, impl2.err
		}
		if impl3.err != nil {
			return Tuple7[T1, T2, T3, T4, T5, T6, T7]{}, impl3.err
		}
		if impl4.err != nil {
			return Tuple7[T1, T2, T3, T4, T5, T6, T7]{}, impl4.err
		}
		if impl5.err != nil {
			return Tuple7[T1, T2, T3, T4, T5, T6, T7]{}, impl5.err
		}
		if impl6.err != nil {
			return Tuple7[T1, T2, T3, T4, T5, T6, T7]{}, impl6.err
		}
		if impl7.err != nil {
			return Tuple7[T1, T2, T3, T4, T5, T6, T7]{}, impl7.err
		}

		return NewTuple7(impl1.value, impl2.value, impl3.value, impl4.value, impl5.value, impl6.value, impl7.value), nil
	})
}

// FutureJoin8 executes eight futures concurrently and combines their results
// into a Tuple8. Returns an error if any future fails.
//
// The returned future respects context cancellation when Get(ctx) is called.
func FutureJoin8[T1, T2, T3, T4, T5, T6, T7, T8 any](f1 Future[T1], f2 Future[T2], f3 Future[T3], f4 Future[T4], f5 Future[T5], f6 Future[T6], f7 Future[T7], f8 Future[T8]) Future[Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]] {
	impl1 := f1.(*future[T1])
	impl2 := f2.(*future[T2])
	impl3 := f3.(*future[T3])
	impl4 := f4.(*future[T4])
	impl5 := f5.(*future[T5])
	impl6 := f6.(*future[T6])
	impl7 := f7.(*future[T7])
	impl8 := f8.(*future[T8])

	return FutureDo(func() (Tuple8[T1, T2, T3, T4, T5, T6, T7, T8], error) {
		var wg sync.WaitGroup
		wg.Add(8)

		go func() { defer wg.Done(); <-impl1.done }()
		go func() { defer wg.Done(); <-impl2.done }()
		go func() { defer wg.Done(); <-impl3.done }()
		go func() { defer wg.Done(); <-impl4.done }()
		go func() { defer wg.Done(); <-impl5.done }()
		go func() { defer wg.Done(); <-impl6.done }()
		go func() { defer wg.Done(); <-impl7.done }()
		go func() { defer wg.Done(); <-impl8.done }()

		wg.Wait()

		if impl1.err != nil {
			return Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]{}, impl1.err
		}
		if impl2.err != nil {
			return Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]{}, impl2.err
		}
		if impl3.err != nil {
			return Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]{}, impl3.err
		}
		if impl4.err != nil {
			return Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]{}, impl4.err
		}
		if impl5.err != nil {
			return Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]{}, impl5.err
		}
		if impl6.err != nil {
			return Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]{}, impl6.err
		}
		if impl7.err != nil {
			return Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]{}, impl7.err
		}
		if impl8.err != nil {
			return Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]{}, impl8.err
		}

		return NewTuple8(impl1.value, impl2.value, impl3.value, impl4.value, impl5.value, impl6.value, impl7.value, impl8.value), nil
	})
}
