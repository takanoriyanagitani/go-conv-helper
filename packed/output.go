package packed

import (
	"context"

	ch "github.com/takanoriyanagitani/go-conv-helper"
)

type PackedSource[P any] func() (packed P, e error)

type UnpackPackedInput[P, I any] func(packed P) (unpacked []I, e error)

type Output2Packed[O, Q any] func(packedOutput Q, output O) (Q, error)

type ConvertPacked[P, I, O, Q any] func(ctx context.Context, ipack P) (opack Q, e error)

type ConvertPackedBuilder[P, I, O, Q any] struct {
	UnpackPackedInput[P, I]
	ch.Converter[I, O]
	Output2Packed[O, Q]
}

func (b ConvertPackedBuilder[P, I, O, Q]) Build() ConvertPacked[P, I, O, Q] {
	return func(ctx context.Context, packedSrc P) (packedOut Q, e error) {
		unpacked, ue := b.UnpackPackedInput(packedSrc)
		if nil != ue {
			return packedOut, ue
		}
		for _, input := range unpacked {
			converted, ce := b.Converter(ctx, input)
			if nil != ce {
				return packedOut, ce
			}
			packedOut, e = b.Output2Packed(packedOut, converted)
			if nil != e {
				return
			}
		}
		return
	}
}
