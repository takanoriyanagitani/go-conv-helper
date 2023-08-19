package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	gc "google.golang.org/grpc/credentials"
	gci "google.golang.org/grpc/credentials/insecure"

	ch "github.com/takanoriyanagitani/go-conv-helper"
	h2r "github.com/takanoriyanagitani/go-conv-helper/examples/external/net/grpc/python/color/proto/color/hsv/v1"
)

type Hsv struct {
	hue        float32
	saturation float32
	value      float32
}

func (h Hsv) toProto() *h2r.Hsv {
	return &h2r.Hsv{
		Hue:        h.hue,
		Saturation: h.saturation,
		Value:      h.value,
	}
}

func (h Hsv) toRequest() *h2r.FromHsvRequest {
	var hsv *h2r.Hsv = h.toProto()
	return &h2r.FromHsvRequest{
		Hsv: hsv,
	}
}

func (r Rgb) fromProto(p *h2r.Rgb) Rgb {
	return Rgb{
		red:   p.GetRed(),
		blue:  p.GetBlue(),
		green: p.GetGreen(),
	}
}

func (r Rgb) fromResponse(p *h2r.FromHsvResponse) Rgb {
	var rgb *h2r.Rgb = p.GetRgb()
	return r.fromProto(rgb)
}

type Rgb struct {
	red   float32
	green float32
	blue  float32
}

func g2conv(client h2r.ColorConvServiceClient) ch.Converter[Hsv, Rgb] {
	return func(ctx context.Context, h Hsv) (Rgb, error) {
		var req *h2r.FromHsvRequest = h.toRequest()
		res, e := client.FromHsv(ctx, req)
		var rgb Rgb = Rgb{}.fromResponse(res)
		return rgb, e
	}
}

const addr string = "localhost:50051"

var tc gc.TransportCredentials = gci.NewCredentials()
var opts []grpc.DialOption = []grpc.DialOption{
	grpc.WithTransportCredentials(tc),
}

func main() {
	conn, e := grpc.Dial(addr, opts...)
	if nil != e {
		log.Fatal(e)
	}
	defer conn.Close()

	var cli h2r.ColorConvServiceClient = h2r.NewColorConvServiceClient(conn)
	var cnv ch.Converter[Hsv, Rgb] = g2conv(cli)
	rgb, e := cnv(context.Background(), Hsv{
		hue:        0.5,
		saturation: 0.5,
		value:      1.0,
	})
	if nil != e {
		log.Fatal(e)
	}
	log.Printf("rgb: %v\n", rgb)
}
