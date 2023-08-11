package ir

import (
	"context"
	"time"
)

type BRPop[K Key] func(ctx context.Context, timeout time.Duration, key K) (string, error)
