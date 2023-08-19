package util

import (
	"context"
)

func NoErr[I, O any](original func(I) O) func(context.Context, I) (O, error) {
	return func(ctx context.Context, input I) (output O, e error) {
		return original(input), nil
	}
}
