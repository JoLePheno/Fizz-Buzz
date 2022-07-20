package port

import "errors"

var (
	ErrInvalidLimit    = errors.New("limit must be bigger than 0")
	ErrInvalidIntegers = errors.New("int1 must be lower than int2")
)
