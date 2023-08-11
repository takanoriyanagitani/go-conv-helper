package gr

import (
	"context"
	"errors"
	"time"

	g9 "github.com/redis/go-redis/v9"

	util "github.com/takanoriyanagitani/go-conv-helper/util"

	ir "github.com/takanoriyanagitani/go-conv-helper/external/buffer/redis"
)

var ErrUnexpectedResult error = errors.New("unexpected result")

type BRPop func(ctx context.Context, timeout time.Duration, key string) *g9.StringSliceCmd

func BRPopNew(client *g9.Client) BRPop {
	return func(ctx context.Context, timeout time.Duration, key string) *g9.StringSliceCmd {
		return client.BRPop(ctx, timeout, key)
	}
}

func (g BRPop) strings2result(input []string) (string, error) {
	switch len(input) {
	case 2:
		return input[1], nil
	default:
		return "", ErrUnexpectedResult
	}
}

func (g BRPop) result2strings(result *g9.StringSliceCmd) ([]string, error) {
	return result.Result()
}

func (g BRPop) ToBRPop() ir.BRPop[string] {
	var result2string func(*g9.StringSliceCmd) (string, error) = util.ComposeErr(
		g.result2strings,
		g.strings2result,
	)
	return func(ctx context.Context, timeout time.Duration, key string) (string, error) {
		var ret *g9.StringSliceCmd = g(ctx, timeout, key)
		return result2string(ret)
	}
}
