package util

import (
	"errors"
)

var ErrOutOfBound error = errors.New("out of bound")

type Array[T any] []T

func (a Array[T]) First() (t T, e error) {
	switch len(a) {
	case 0:
		return t, ErrOutOfBound
	default:
		return a[0], nil
	}
}
