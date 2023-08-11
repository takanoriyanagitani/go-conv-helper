package ir

import (
	"context"
	"errors"

	util "github.com/takanoriyanagitani/go-conv-helper/util"
)

var ErrTooManyRequests error = errors.New("too many requests")

type LengthChecker func(length int64) error

type SimpleResponseKeyBuilder func(requestKey string) (responseKey string)

func SimpleResponseKeyBuilderNewStatic(suffix string) SimpleResponseKeyBuilder {
	return func(requestKey string) (responseKey string) {
		return requestKey + suffix
	}
}

var SimpleResponseKeyBuilderStaticDefault SimpleResponseKeyBuilder = SimpleResponseKeyBuilderNewStatic("-response")

func LengthCheckerNewStatic(upperBound int64) LengthChecker {
	return func(length int64) error {
		var tooMany bool = upperBound <= length
		return util.Select(
			nil,
			ErrTooManyRequests,
			tooMany,
		)
	}
}

var LengthChecker1K LengthChecker = LengthCheckerNewStatic(1000)
var LengthCheckerDefault LengthChecker = LengthChecker1K

type LengthLimiter func(ctx context.Context) error

type LengthGetter[K Key] LLen[K]

func (g LengthGetter[K]) NewLimiter(checker LengthChecker) func(K) LengthLimiter {
	return func(key K) LengthLimiter {
		return func(ctx context.Context) error {
			length, e := g(ctx, key)
			return util.Select(
				e,
				checker(length),
				nil == e,
			)
		}
	}
}

type Serialize[I Input, V Val] func(input I) (serialized V, e error)
type Parse[O Output, V Val] func(val V) (parsed O, e error)

type ParseFromString[O Output] Parse[O, string]
type SerializeToBytes[I Input] Serialize[I, []byte]

type PushInput[I Input] func(ctx context.Context, input I) (length int64, e error)

type PushInputBuilder[I Input, K Key, V Val] struct {
	Serialize[I, V]
	LPush[K, V]
	Key K
}

func (b PushInputBuilder[I, K, V]) serialize(input I) (serialized V, e error) {
	return b.Serialize(input)
}

func (b PushInputBuilder[I, K, V]) lpush(ctx context.Context, val V) (length int64, e error) {
	return b.LPush(ctx, b.Key, val)
}

func (b PushInputBuilder[I, K, V]) Build() PushInput[I] {
	return util.ComposeCtx(
		func(_ context.Context, input I) (serialized V, e error) { return b.serialize(input) },
		b.lpush,
	)
}
