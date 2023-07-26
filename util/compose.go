package util

import (
	"context"
)

func ComposeErr[T, U, V any](f func(T) (U, error), g func(U) (V, error)) func(T) (V, error) {
	return func(t T) (v V, e error) {
		u, e := f(t)
		return Select(
			func() (V, error) { return v, e },
			func() (V, error) { return g(u) },
			nil == e,
		)()
	}
}

func ComposeCtx[T, U, V any](
	f func(context.Context, T) (U, error),
	g func(context.Context, U) (V, error),
) func(context.Context, T) (V, error) {
	return func(ctx context.Context, t T) (v V, e error) {
		u, e := f(ctx, t)
		return Select(
			func() (V, error) { return v, e },
			func() (V, error) { return g(ctx, u) },
			nil == e,
		)()
	}
}
