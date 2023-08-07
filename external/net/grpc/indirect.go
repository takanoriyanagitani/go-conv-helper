package gconv

import (
	"context"

	"google.golang.org/grpc"

	ch "github.com/takanoriyanagitani/go-conv-helper"
	util "github.com/takanoriyanagitani/go-conv-helper/util"

	raw "github.com/takanoriyanagitani/go-conv-helper/external/net/grpc/proto/raw/v1"
)

type ClientToConverter func(raw.ConvertServiceClient) ch.RawConverter

type ClientToConverterBuilder struct {
	RequestCreator
	ResponseSerializer
	Options []grpc.CallOption
}

func (b ClientToConverterBuilder) Input2Request(input []byte) (*raw.ConvertRequest, error) {
	return b.RequestCreator(input)
}

func (b ClientToConverterBuilder) Response2Output(res *raw.ConvertResponse) (output []byte, e error) {
	return b.ResponseSerializer(res)
}

func (b ClientToConverterBuilder) Request2Response(
	ctx context.Context,
	req *raw.ConvertRequest,
	cli raw.ConvertServiceClient,
) (*raw.ConvertResponse, error) {
	return cli.Convert(ctx, req, b.Options...)
}

func (b ClientToConverterBuilder) ToInput2Response(cli raw.ConvertServiceClient) Input2Response {
	return util.ComposeCtx(
		func(_ context.Context, input []byte) (*raw.ConvertRequest, error) {
			return b.Input2Request(input)
		},
		func(ctx context.Context, req *raw.ConvertRequest) (*raw.ConvertResponse, error) {
			return b.Request2Response(ctx, req, cli)
		},
	)
}

func (b ClientToConverterBuilder) Build() ClientToConverter {
	return func(cli raw.ConvertServiceClient) ch.RawConverter {
		var i2r Input2Response = b.ToInput2Response(cli)
		return b.ResponseSerializer.ToConverter(i2r)
	}
}

func (b ClientToConverterBuilder) WithRequestCreator(rc RequestCreator) ClientToConverterBuilder {
	return ClientToConverterBuilder{
		RequestCreator:     rc,
		ResponseSerializer: b.ResponseSerializer,
		Options:            b.Options,
	}
}

func (b ClientToConverterBuilder) WithSerializer(rs ResponseSerializer) ClientToConverterBuilder {
	return ClientToConverterBuilder{
		RequestCreator:     b.RequestCreator,
		ResponseSerializer: rs,
		Options:            b.Options,
	}
}

func (b ClientToConverterBuilder) WithOptions(opts []grpc.CallOption) ClientToConverterBuilder {
	return ClientToConverterBuilder{
		RequestCreator:     b.RequestCreator,
		ResponseSerializer: b.ResponseSerializer,
		Options:            opts,
	}
}

var ClientToConverterBuilderEmpty ClientToConverterBuilder
var ClientToConverterBuilderDefault ClientToConverterBuilder = ClientToConverterBuilderEmpty.
	WithRequestCreator(DefaultRequestCreator).
	WithSerializer(DefaultResponseSerializer).
	WithOptions(nil)

var ClientToConverterDefault ClientToConverter = ClientToConverterBuilderDefault.Build()
