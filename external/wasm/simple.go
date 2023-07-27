package wc

import (
	"context"
)

type SimpleFloat interface {
	float32 | float64
}

type SimplePositive interface {
	uint32 | uint64
}

type SimpleInteger interface {
	int32 | int64
}

type SimpleType interface {
	SimpleFloat | SimplePositive | SimpleInteger
}

type SimpleFunc[I, O SimpleType] func(ctx context.Context, input I) (output O, e error)

type SimpleFuncII[T, U, V SimpleType] func(ctx context.Context, t T, u U) (v V, e error)

type SimpleFuncBuilder[I, O SimpleType] func(ctx context.Context, wasm []byte) (SimpleFunc[I, O], error)

type SimpleFuncBuilderII[T, U, V SimpleType] func(
	ctx context.Context,
	wasm []byte,
) (SimpleFuncII[T, U, V], error)
