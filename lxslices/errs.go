package lxslices

import "errors"

var (
	ErrDuplicateKey = errors.New("lxslices: duplicate key")
	ErrInvalidSize  = errors.New("lxslices: size must be greater than 0")
)
