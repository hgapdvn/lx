package lxtypes

// Predicate represents a function that tests a condition on an input value.
// Returns true if the input matches the condition, false otherwise.
type Predicate[T any] func(T) bool

// And returns a composed predicate that represents a short-circuiting logical
// AND of this predicate and another.
func (p Predicate[T]) And(other Predicate[T]) Predicate[T] {
	return func(t T) bool {
		return p(t) && other(t)
	}
}

// Or returns a composed predicate that represents a short-circuiting logical
// OR of this predicate and another.
func (p Predicate[T]) Or(other Predicate[T]) Predicate[T] {
	return func(t T) bool {
		return p(t) || other(t)
	}
}

// Negate returns a predicate that represents the logical negation of this predicate.
func (p Predicate[T]) Negate() Predicate[T] {
	return func(t T) bool {
		return !p(t)
	}
}

// BiPredicate represents a function that tests a condition on two input values.
type BiPredicate[T, U any] func(T, U) bool

// And returns a composed predicate that represents a short-circuiting logical
// AND of this predicate and another.
func (bp BiPredicate[T, U]) And(other BiPredicate[T, U]) BiPredicate[T, U] {
	return func(t T, u U) bool {
		return bp(t, u) && other(t, u)
	}
}

// Or returns a composed predicate that represents a short-circuiting logical
// OR of this predicate and another.
func (bp BiPredicate[T, U]) Or(other BiPredicate[T, U]) BiPredicate[T, U] {
	return func(t T, u U) bool {
		return bp(t, u) || other(t, u)
	}
}

// Negate returns a predicate that represents the logical negation of this predicate.
func (bp BiPredicate[T, U]) Negate() BiPredicate[T, U] {
	return func(t T, u U) bool {
		return !bp(t, u)
	}
}

// Consumer represents an operation that accepts a single input argument and returns no result.
type Consumer[T any] func(T)

// AndThen returns a composed Consumer that performs this operation followed by the after operation.
func (c Consumer[T]) AndThen(after Consumer[T]) Consumer[T] {
	return func(t T) {
		c(t)
		after(t)
	}
}

// BiConsumer represents an operation that accepts two input arguments and returns no result.
type BiConsumer[T, U any] func(T, U)

// AndThen returns a composed BiConsumer that performs this operation followed by the after operation.
func (bc BiConsumer[T, U]) AndThen(after BiConsumer[T, U]) BiConsumer[T, U] {
	return func(t T, u U) {
		bc(t, u)
		after(t, u)
	}
}

// Function represents a function that accepts one argument and produces a result.
type Function[T, U any] func(T) U

// AndThen returns a composed function that first applies this function to its input,
// and then applies the after function to the result.
// f.AndThen(g)(x) = g(f(x))
//
// Note: Due to Go's limitations, the after function must accept and return the same type U.
// For chaining functions with different output types, use the standalone Compose function.
func (f Function[T, U]) AndThen(after func(U) U) Function[T, U] {
	return func(t T) U {
		return after(f(t))
	}
}

// Compose returns a composed function that first applies the before function to its input,
// and then applies this function to the result.
// f.Compose(g)(x) = f(g(x))
//
// Note: Due to Go's limitations, the before function must accept and return the same type T.
// For chaining functions with different input types, use the standalone Compose function.
func (f Function[T, U]) Compose(before func(T) T) Function[T, U] {
	return func(t T) U {
		return f(before(t))
	}
}

// BiFunction represents a function that accepts two arguments and produces a result.
type BiFunction[T, U, R any] func(T, U) R

// AndThen returns a composed function that first applies this function to its inputs,
// and then applies the after function to the result.
// bf.AndThen(g)(x, y) = g(bf(x, y))
//
// Note: Due to Go's limitations, the after function must accept and return the same type R.
// For transforming to different types, use the standalone Compose function with the result.
func (bf BiFunction[T, U, R]) AndThen(after func(R) R) BiFunction[T, U, R] {
	return func(t T, u U) R {
		return after(bf(t, u))
	}
}

// Supplier represents a supplier of results. Each call may return a different result.
type Supplier[T any] func() T

// UnaryOperator represents an operation on a single operand that produces a result
// of the same type as its operand.
type UnaryOperator[T any] func(T) T

// BinaryOperator represents an operation upon two operands of the same type,
// producing a result of the same type as the operands.
type BinaryOperator[T any] func(T, T) T

// Comparator represents a comparison function which imposes a total ordering on some
// collection of objects. Returns a negative integer, zero, or a positive integer
// as the first argument is less than, equal to, or greater than the second.
type Comparator[T any] func(T, T) int

// Reversed returns a comparator that imposes the reverse ordering of this comparator.
func (c Comparator[T]) Reversed() Comparator[T] {
	return func(t1, t2 T) int {
		return c(t2, t1)
	}
}

// ThenComparing returns a lexicographic-order comparator with another comparator.
func (c Comparator[T]) ThenComparing(other Comparator[T]) Comparator[T] {
	return func(t1, t2 T) int {
		res := c(t1, t2)
		if res != 0 {
			return res
		}
		return other(t1, t2)
	}
}
