package gconv

import (
	"context"

	ch "github.com/takanoriyanagitani/go-conv-helper"
	util "github.com/takanoriyanagitani/go-conv-helper/util"

	raw "github.com/takanoriyanagitani/go-conv-helper/external/net/grpc/proto/raw/v1"
)

type RawConverterBuilder func(raw.ConvertServiceServer) ch.RawConverter

type RequestCreator func(input []byte) (*raw.ConvertRequest, error)
type Input2Response func(ctx context.Context, input []byte) (*raw.ConvertResponse, error)

func (c RequestCreator) ToInput2Response(server raw.ConvertServiceServer) Input2Response {
	return util.ComposeCtx(
		func(_ context.Context, input []byte) (*raw.ConvertRequest, error) { return c(input) },
		server.Convert,
	)
}

func RequestCreatorNew(typename string) RequestCreator {
	return func(input []byte) (*raw.ConvertRequest, error) {
		return &raw.ConvertRequest{
			Type: &raw.Type{Name: typename},
			Data: &raw.Data{Raw: input},
		}, nil
	}
}

var DefaultRequestCreator RequestCreator = RequestCreatorNew("")

type ResponseSerializer func(*raw.ConvertResponse) (output []byte, e error)

func ResponseSerializerNew(
	serializer func(*raw.Type, *raw.Data) ([]byte, error),
) ResponseSerializer {
	return func(res *raw.ConvertResponse) (output []byte, e error) {
		return serializer(res.Type, res.Data)
	}
}

var DefaultResponseSerializer ResponseSerializer = ResponseSerializerNew(
	func(_ *raw.Type, data *raw.Data) (output []byte, e error) { return data.Raw, nil },
)

func RawConverterBuilderNew(q2i RequestCreator) func(ResponseSerializer) RawConverterBuilder {
	return func(ser ResponseSerializer) RawConverterBuilder {
		return func(server raw.ConvertServiceServer) ch.RawConverter {
			var i2r Input2Response = q2i.ToInput2Response(server)
			return util.ComposeCtx(
				i2r,
				func(_ context.Context, res *raw.ConvertResponse) ([]byte, error) { return ser(res) },
			)
		}
	}
}

var DefaultConverterBuilder RawConverterBuilder = RawConverterBuilderNew(
	DefaultRequestCreator,
)(DefaultResponseSerializer)
