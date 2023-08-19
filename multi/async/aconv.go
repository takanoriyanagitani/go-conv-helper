package aconv

import (
	"context"

	ch "github.com/takanoriyanagitani/go-conv-helper"
)

type Either[V any] struct {
	left  error
	right V
}

func (e Either[V]) Left() error { return e.left }
func (e Either[V]) Right() V    { return e.right }

type ConvertToChan[I, O any] func(ctx context.Context, input I, out chan<- Either[O])

func ConvertToChanNew[I, O any](conv ch.Converter[I, O]) ConvertToChan[I, O] {
	return func(ctx context.Context, input I, out chan<- Either[O]) {
		go func() {
			o, e := conv(ctx, input)
			var oe Either[O] = Either[O]{
				left:  e,
				right: o,
			}
			out <- oe
		}()
	}
}

func ConverterMultiNew[I, V, O any](
	reducer func(state O, next Either[V]) (O, error),
	converters ...ConvertToChan[I, V],
) ch.Converter[I, O] {
	return func(ctx context.Context, input I) (output O, e error) {
		var sz int = len(converters)
		var c chan Either[V] = make(chan Either[V], sz)
		for _, conv := range converters {
			conv(ctx, input, c)
		}
		for i := 0; i < sz; i++ {
			var result Either[V] = <-c
			output, e = reducer(output, result)
			if nil != e {
				return output, e
			}
		}
		close(c)
		return output, e
	}
}
