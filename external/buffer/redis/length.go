package ir

import (
	"context"
)

type LLen[K Key] func(context.Context, K) (length int64, e error)
