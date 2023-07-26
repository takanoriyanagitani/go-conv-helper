package hconv

import (
	"bytes"
	"context"
	"io"
	"net/http"

	ch "github.com/takanoriyanagitani/go-conv-helper"
	util "github.com/takanoriyanagitani/go-conv-helper/util"
)

type Builder[I, O any] struct {
	Method string
	Url    string
	ConvI  func(I) ([]byte, error)
	ConvO  func(*http.Response) (O, error)
	Client *http.Client
}

func (b Builder[I, O]) NewInput2Request() func(context.Context, I) (*http.Request, error) {
	return func(ctx context.Context, input I) (*http.Request, error) {
		return util.ComposeErr(
			b.ConvI,
			func(body []byte) (*http.Request, error) {
				var reader io.Reader = bytes.NewReader(body)
				return http.NewRequestWithContext(
					ctx,
					b.Method,
					b.Url,
					reader,
				)
			},
		)(input)
	}
}

func (b Builder[I, O]) NewRequest2Output() func(*http.Request) (O, error) {
	return util.ComposeErr(
		b.Client.Do,
		b.ConvO,
	)
}

func (b Builder[I, O]) NewInput2Output() func(context.Context, I) (O, error) {
	var i2r func(context.Context, I) (*http.Request, error) = b.NewInput2Request()
	var r2o func(*http.Request) (O, error) = b.NewRequest2Output()
	return util.ComposeCtx(
		i2r,
		func(_ context.Context, req *http.Request) (O, error) { return r2o(req) },
	)
}

func (b Builder[I, O]) Build() ch.Converter[I, O] { return b.NewInput2Output() }
