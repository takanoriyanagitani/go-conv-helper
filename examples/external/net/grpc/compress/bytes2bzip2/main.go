package main

import (
	"bytes"
	"compress/bzip2"
	"context"
	"io"
	"log"
	"strings"

	"google.golang.org/grpc"
	gc "google.golang.org/grpc/credentials"
	gci "google.golang.org/grpc/credentials/insecure"

	ch "github.com/takanoriyanagitani/go-conv-helper"
	gconv "github.com/takanoriyanagitani/go-conv-helper/external/net/grpc"
	raw "github.com/takanoriyanagitani/go-conv-helper/external/net/grpc/proto/raw/v1"

	b2b "github.com/takanoriyanagitani/go-conv-helper/examples/external/net/grpc/compress/bytes2bzip2/proto/bytes2bz/v1"
)

const addr string = "localhost:50051"

type ConvertServiceClientFunc func(context.Context, *raw.ConvertRequest) (*raw.ConvertResponse, error)

func (f ConvertServiceClientFunc) Convert(
	ctx context.Context,
	req *raw.ConvertRequest,
	opts ...grpc.CallOption,
) (*raw.ConvertResponse, error) {
	return f(ctx, req)
}

func ConvertServiceClientFuncNew(cli b2b.CompressServiceClient) ConvertServiceClientFunc {
	return func(ctx context.Context, rawq *raw.ConvertRequest) (*raw.ConvertResponse, error) {
		var data *raw.Data = rawq.Data
		var input []byte = data.Raw
		var req *b2b.CompressRequest = &b2b.CompressRequest{
			Level: b2b.Compression_LEVEL_NUM_FAST,
			Input: input,
		}
		var res *b2b.CompressResponse = must(cli.Compress(ctx, req))
		var typ *raw.Type = &raw.Type{Name: "application/x-bzip2"}
		return &raw.ConvertResponse{
			Type: typ,
			Data: &raw.Data{Raw: res.Compressed},
		}, nil
	}
}

var h2c gc.TransportCredentials = gci.NewCredentials()
var opts []grpc.DialOption = []grpc.DialOption{
	grpc.WithTransportCredentials(h2c),
}

func must[T any](t T, e error) T {
	switch e {
	case nil:
		return t
	default:
		panic(e)
	}
}

func main() {
	var conn *grpc.ClientConn = must(grpc.Dial(addr, opts...))
	defer conn.Close()
	var cli b2b.CompressServiceClient = b2b.NewCompressServiceClient(conn)
	var cvc raw.ConvertServiceClient = ConvertServiceClientFuncNew(cli)
	var c2c gconv.ClientToConverter = gconv.ClientToConverterDefault
	var cnv ch.RawConverter = c2c(cvc)
	const input string = "hw"
	var ctx context.Context = context.Background()
	var compressed []byte = must(cnv(ctx, []byte(input)))
	var rdr *bytes.Reader = bytes.NewReader(compressed)
	var brdr io.Reader = bzip2.NewReader(rdr)
	var buf strings.Builder
	_, e := io.Copy(&buf, brdr)
	if nil != e {
		log.Fatal(e)
	}
	log.Printf("original: %s\n", input)
	log.Printf("decompressed: %s\n", buf.String())
}
