package sconv

import (
	"context"

	ch "github.com/takanoriyanagitani/go-conv-helper"
)

func ConverterMultiNew[I, V, O any](
	finalizer ch.Converter[[]V, O],
	converters ...ch.Converter[I, V],
) ch.Converter[I, O] {
	return func(ctx context.Context, input I) (output O, e error) {
		var results []V
		for _, conv := range converters {
			v, e := conv(ctx, input)
			if nil != e {
				return output, e
			}
			results = append(results, v)
		}
		return finalizer(ctx, results)
	}
}
