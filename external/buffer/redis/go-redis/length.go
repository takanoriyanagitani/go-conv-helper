package gr

import (
	"context"

	g9 "github.com/redis/go-redis/v9"

	ir "github.com/takanoriyanagitani/go-conv-helper/external/buffer/redis"
)

type LLen func(ctx context.Context, key string) *g9.IntCmd

func (g LLen) ToLLen() ir.LLen[string] {
	return func(ctx context.Context, key string) (length int64, e error) {
		var ret *g9.IntCmd = g(ctx, key)
		return ret.Result()
	}
}

func LLenNew(client *g9.Client) LLen {
	return func(ctx context.Context, key string) *g9.IntCmd {
		return client.LLen(ctx, key)
	}
}
