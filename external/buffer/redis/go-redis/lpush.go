package gr

import (
	"context"

	g9 "github.com/redis/go-redis/v9"

	ir "github.com/takanoriyanagitani/go-conv-helper/external/buffer/redis"
)

type LPush[V ir.Val] func(ctx context.Context, key string, val V) *g9.IntCmd

func LPushNew[V ir.Val](client *g9.Client) LPush[V] {
	return func(ctx context.Context, key string, val V) *g9.IntCmd {
		return client.LPush(ctx, key, val)
	}
}

func (g LPush[V]) ToLPush() ir.LPush[string, V] {
	return func(ctx context.Context, key string, val V) (length int64, e error) {
		var ret *g9.IntCmd = g(ctx, key, val)
		return ret.Result()
	}
}
