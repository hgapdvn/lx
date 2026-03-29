// Package lxtypes provides reusable generic type definitions for functional programming,
// optional values, error handling, and asynchronous operations in Go.
//
// This package includes four main categories of types:
//
// 1. Functional Interfaces (inspired by Java's java.util.function):
//
//   - Predicate[T], BiPredicate[T, U] - Test conditions
//   - Consumer[T], BiConsumer[T, U] - Perform operations
//   - Function[T, U], BiFunction[T, U, R] - Transform values
//   - Supplier[T] - Provide values
//   - UnaryOperator[T], BinaryOperator[T] - Operate on same type
//   - Comparator[T] - Compare values for ordering
//
// 2. Optional and Error Handling:
//
//   - Optional[T] - Optional value (Java-style with comma-ok pattern: Get() returns (T, bool))
//   - Result[T] - Error handling with Go's (value, error) pattern (Value() returns (T, error))
//   - Either[L, R] - General binary choice between any two types (EitherLeft, EitherRight)
//
// 3. Tuple Types:
//
//   - Pair[T, U], Triple[T, U, V], Quad[T, U, V, W] - Multi-value tuples for 2-4 values
//   - Tuple5[...], Tuple6[...], Tuple7[...], Tuple8[...] - Extended tuples for 5-8 values
//
// 4. Lazy Evaluation:
//
//   - Lazy[T] - Deferred or immediate computation with caching
//
// 5. Async Operations:
//
//   - Future[T] - Asynchronous computation with type-safe composition and context support
//   - FutureAny[T] - Return the first successful result from many futures (first err==nil)
//
// 6. Mutable State:
//
//   - Ref[T] - Thread-safe mutable value cell (Get, Set, Update)
//
// Quick Examples:
//
//	// Functional types
//	isEven := lxtypes.Predicate[int](func(n int) bool { return n%2 == 0 })
//	fmt.Println(isEven(4))  // true
//
//	// Function composition
//	double := lxtypes.Function[int, int](func(n int) int { return n * 2 })
//	result := double.AndThen(func(n int) int { return n + 10 })(5)  // 20
//
//	// BiFunction composition
//	add := lxtypes.BiFunction[int, int, int](func(a, b int) int { return a + b })
//	result := add.AndThen(func(n int) int { return n * 2 })(3, 4)  // 14
//
//	// Optional values with comma-ok pattern (idiomatic Go)
//	opt := lxtypes.OptionalOf(42)
//	if value, ok := opt.Get(); ok {
//	    fmt.Println(value)  // 42
//	}
//
//	// Or use default values
//	value := opt.OrElse(0)  // 42
//
//	// Safe nil handling with OptionalOfNullable
//	var ptr *int
//	opt2 := lxtypes.OptionalOfNullable(ptr)  // Empty Optional
//	value2 := opt2.OrElse(99)                // 99
//
//	// Error handling with Result[T] using (value, error) pattern
//	result := lxtypes.ResultSuccess(42)
//	if value, err := result.Value(); err == nil {
//	    fmt.Println(value)  // 42
//	}
//
//	// Or use default value
//	failure := lxtypes.ResultFailure[int](errors.New("error"))
//	value := failure.ValueOr(99)  // 99
//
//	// General binary choice with Either[L, R]
//	either := lxtypes.EitherRight[string, int](42)
//	if right, ok := either.Right(); ok {
//	    fmt.Println(right)  // 42
//	}
//
//	// Tuples
//	p := lxtypes.NewPair(42, "answer")
//	fmt.Println(p.First, p.Second)  // 42 answer
//
//	// Extended tuples for 5-8 values
//	t5 := lxtypes.NewTuple5(1, "two", true, 4.0, []int{5, 6})
//	fmt.Println(t5.V1, t5.V2, t5.V3, t5.V4, t5.V5)  // 1 two true 4.0 [5 6]
//
//	// Lazy evaluation - deferred computation
//	expensive := lxtypes.LazyDeferred(func() (int, error) {
//	    // Expensive computation only runs when needed
//	    return 42, nil
//	})
//	value, _ := expensive.Get()  // Computed here
//	value2, _ := expensive.Get() // Returns cached value
//
//	// Lazy evaluation - immediate value
//	immediate := lxtypes.LazyEager(100)
//	value, _ := immediate.Get()  // Returns immediately
//
//	// Async operations - parallel execution
//	f1 := lxtypes.FutureDo(func() (int, error) { return fetchData1() })
//	f2 := lxtypes.FutureDo(func() (int, error) { return fetchData2() })
//	allData := lxtypes.FutureAll(f1, f2)
//	results, _ := allData.Get(context.Background())  // [data1, data2]
//
//	// Async operations - FutureAny (first successful result)
//	fast := lxtypes.FutureDo(func() (string, error) { return "fast", nil })
//	slow := lxtypes.FutureDo(func() (string, error) { time.Sleep(50 * time.Millisecond); return "slow", nil })
//	any := lxtypes.FutureAny(fast, slow)
//	val, _ := any.Get(context.Background()) // "fast"
//
//	// Async operations - sequential with type transformation
//	future := lxtypes.FutureDo(func() (int, error) { return getUserId() })
//	future = lxtypes.FutureThen(future, func(id int) (User, error) { return fetchUser(id) })
//	future = lxtypes.FutureThen(future, func(user User) (Account, error) { return fetchAccount(user) })
//	account, _ := future.Get(context.Background())
//
//	// Async operations - join different types
//	userFuture := lxtypes.FutureDo(func() (User, error) { return fetchUser() })
//	configFuture := lxtypes.FutureDo(func() (Config, error) { return fetchConfig() })
//	combined := lxtypes.FutureJoin2(userFuture, configFuture)
//	result, _ := combined.Get(context.Background())  // Pair[User, Config]
//
//	// Async operations - join 5+ different types (microservices)
//	user := lxtypes.FutureDo(func() (User, error) { return fetchUser() })
//	orders := lxtypes.FutureDo(func() ([]Order, error) { return fetchOrders() })
//	payment := lxtypes.FutureDo(func() (Payment, error) { return fetchPayment() })
//	inventory := lxtypes.FutureDo(func() (Inventory, error) { return fetchInventory() })
//	recommendations := lxtypes.FutureDo(func() ([]Product, error) { return fetchRecommendations() })
//	all := lxtypes.FutureJoin5(user, orders, payment, inventory, recommendations)
//	dashboard, _ := all.Get(context.Background())  // Tuple5[User, []Order, Payment, Inventory, []Product]
//
//	// Mutable state - thread-safe value cell
//	counter := lxtypes.NewRef(0)
//	counter.Update(func(v int) int { return v + 1 })
//	fmt.Println(counter.Get())  // 1
//
//	// Safe for concurrent use
//	var wg sync.WaitGroup
//	for i := 0; i < 10; i++ {
//	    wg.Add(1)
//	    go func() {
//	        defer wg.Done()
//	        counter.Update(func(v int) int { return v + 1 })
//	    }()
//	}
//	wg.Wait()
//	fmt.Println(counter.Get())  // 11
//
// For comprehensive documentation, examples, and use cases, see:
// https://github.com/hgapdvn/lx/tree/main/lxtypes#readme
package lxtypes
