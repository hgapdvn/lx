package lxtypes

// Either represents a value that can be either Left or Right.
type Either[L, R any] interface {
	// Left returns the Left value and true if this is Left, or zero value and false if Right.
	Left() (L, bool)

	// Right returns the Right value and true if this is Right, or zero value and false if Left.
	Right() (R, bool)

	// LeftOr returns the Left value or the provided default if Right.
	LeftOr(defaultValue L) L

	// RightOr returns the Right value or the provided default if Left.
	RightOr(defaultValue R) R
}

// EitherLeft creates an Either with Left value.
func EitherLeft[L, R any](value L) Either[L, R] {
	return eitherLeft[L, R]{value: value}
}

// EitherRight creates an Either with Right value.
func EitherRight[L, R any](value R) Either[L, R] {
	return eitherRight[L, R]{value: value}
}

// ------------------------------------ Either Left implementation ------------------------------------

type eitherLeft[L, R any] struct {
	value L
}

func (e eitherLeft[L, R]) Left() (L, bool) {
	return e.value, true
}

func (e eitherLeft[L, R]) Right() (R, bool) {
	var zero R
	return zero, false
}

func (e eitherLeft[L, R]) LeftOr(defaultValue L) L {
	return e.value
}

func (e eitherLeft[L, R]) RightOr(defaultValue R) R {
	return defaultValue
}

// ------------------------------------ Either Right implementation ------------------------------------

type eitherRight[L, R any] struct {
	value R
}

func (e eitherRight[L, R]) Left() (L, bool) {
	var zero L
	return zero, false
}

func (e eitherRight[L, R]) Right() (R, bool) {
	return e.value, true
}

func (e eitherRight[L, R]) LeftOr(defaultValue L) L {
	return defaultValue
}

func (e eitherRight[L, R]) RightOr(defaultValue R) R {
	return e.value
}
