package conv

import (
	"context"

	util "github.com/takanoriyanagitani/go-conv-helper/util"
)

type RawConverter Converter[[]byte, []byte]

type Serializer[I any] func(input I) (serialized []byte, e error)

type Parser[O any] func(raw []byte) (parsed O, e error)

type ConverterBuilder[I, O any] struct {
	Serializer[I]
	Parser[O]
	RawConverter
}

func (b ConverterBuilder[I, O]) input2rawNew() func(context.Context, I) (raw []byte, e error) {
	return util.ComposeCtx(
		func(_ context.Context, input I) (rawInput []byte, e error) { return b.Serializer(input) },
		b.RawConverter,
	)
}

func (b ConverterBuilder[I, O]) Build() Converter[I, O] {
	return util.ComposeCtx(
		b.input2rawNew(),
		func(_ context.Context, rawOutput []byte) (O, error) { return b.Parser(rawOutput) },
	)
}
