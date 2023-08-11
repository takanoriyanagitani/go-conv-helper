package ir

import (
	"context"
	"time"

	ch "github.com/takanoriyanagitani/go-conv-helper"
	util "github.com/takanoriyanagitani/go-conv-helper/util"
)

type Input any
type Output any

type ConverterBuilder[I Input, O Output, K Key, V Val] struct {
	Serialize[I, V]
	LPush[K, V]
	BRPop[K]
	Parse[O, string]
	RequestKey  K
	ResponseKey K
}

func (c ConverterBuilder[I, O, K, V]) WithSerializer(
	s Serialize[I, V],
) ConverterBuilder[I, O, K, V] {
	c.Serialize = s
	return c
}

func (c ConverterBuilder[I, O, K, V]) WithLPush(l LPush[K, V]) ConverterBuilder[I, O, K, V] {
	c.LPush = l
	return c
}

func (c ConverterBuilder[I, O, K, V]) WithBRPop(b BRPop[K]) ConverterBuilder[I, O, K, V] {
	c.BRPop = b
	return c
}

func (c ConverterBuilder[I, O, K, V]) WithParse(p Parse[O, string]) ConverterBuilder[I, O, K, V] {
	c.Parse = p
	return c
}

func (c ConverterBuilder[I, O, K, V]) WithRequestKey(key K) ConverterBuilder[I, O, K, V] {
	c.RequestKey = key
	return c
}

func (c ConverterBuilder[I, O, K, V]) WithResponseKey(key K) ConverterBuilder[I, O, K, V] {
	c.ResponseKey = key
	return c
}

func (c ConverterBuilder[I, O, K, V]) serialize(input I) (serialized V, e error) {
	return c.Serialize(input)
}

func (c ConverterBuilder[I, O, K, V]) lpush(ctx context.Context, val V) (int64, error) {
	return c.LPush(ctx, c.RequestKey, val)
}

func (c ConverterBuilder[I, O, K, V]) brpop(
	ctx context.Context,
	timeout time.Duration,
) (string, error) {
	return c.BRPop(ctx, timeout, c.ResponseKey)
}

func (c ConverterBuilder[I, O, K, V]) parse(val string) (parsed O, e error) { return c.Parse(val) }

func (c ConverterBuilder[I, O, K, V]) createPush() func(
	ctx context.Context,
	input I,
) (length int64, e error) {
	return util.ComposeCtx(
		func(_ context.Context, input I) (serialized V, e error) { return c.serialize(input) },
		c.lpush,
	)
}

func (c ConverterBuilder[I, O, K, V]) createPop() func(
	ctx context.Context,
	timeout time.Duration,
) (parsed O, e error) {
	return util.ComposeCtx(
		c.brpop,
		func(_ context.Context, val string) (parsed O, e error) { return c.parse(val) },
	)
}

func (c ConverterBuilder[I, O, K, V]) Build(timeout time.Duration) ch.Converter[I, O] {
	var pop func(context.Context, time.Duration) (O, error) = c.createPop()
	return util.ComposeCtx(
		c.createPush(),
		func(ctx context.Context, _ int64) (O, error) { return pop(ctx, timeout) },
	)
}
