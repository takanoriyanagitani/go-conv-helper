package conv

import (
	"context"
)

type InputStore[I, K any] func(ctx context.Context, key K) (input I, e error)
