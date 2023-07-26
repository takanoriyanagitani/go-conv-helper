package conv

import (
	"context"
)

type Converter[I, O any] func(ctx context.Context, input I) (output O, e error)
