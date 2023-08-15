package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	cbor "github.com/fxamacker/cbor/v2"

	ch "github.com/takanoriyanagitani/go-conv-helper"
)

type Timestamp uint64

type Uuid [2]uint64

type ConvertRequest struct {
	Unixtime   Timestamp `json:"timestamp"`
	TraceId    Uuid      `json:"trace_id"`
	Severity   uint8     `json:"severity"`
	Body       string    `json:"body"`
	Name       string    `json:"name"`
	NameSpace  string    `json:"ns"`
	InstanceId Uuid      `json:"instance_id"`
	Version    string    `json:"ver"`
}

type RequestSer func(*ConvertRequest) ([]byte, error)

var CborSer RequestSer = func(req *ConvertRequest) ([]byte, error) { return cbor.Marshal(req) }

type Poster func(contentType string, body io.Reader) (*http.Response, error)

const URL = "http://localhost:8148"

var DefaultPoster Poster = func(ctyp string, bdy io.Reader) (*http.Response, error) {
	var cli *http.Client = http.DefaultClient
	return cli.Post(URL, ctyp, bdy)
}

func PostRequest(poster Poster, req *ConvertRequest, ser RequestSer) (cnv string, e error) {
	serialized, e1 := ser(req)
	res, e2 := poster("application/octet-stream", bytes.NewReader(serialized))
	e12 := errors.Join(e1, e2)
	if nil != e2 {
		return "", e12
	}
	defer res.Body.Close()
	switch res.StatusCode {
	case 200:
		var buf strings.Builder
		_, e3 := io.Copy(&buf, res.Body)
		return buf.String(), errors.Join(e12, e3)
	default:
		return "", errors.Join(e12, fmt.Errorf("Unable to post: %v", res.StatusCode))
	}
}

func NewConverter(ser RequestSer) func(Poster) ch.Converter[*ConvertRequest, string] {
	return func(poster Poster) ch.Converter[*ConvertRequest, string] {
		return func(_ context.Context, input *ConvertRequest) (output string, e error) {
			return PostRequest(poster, input, ser)
		}
	}
}

var ConverterDefault ch.Converter[*ConvertRequest, string] = NewConverter(CborSer)(DefaultPoster)

func main() {
	converted, e := ConverterDefault(
		context.Background(),
		&ConvertRequest{
			Unixtime:   0x1243,
			TraceId:    Uuid{0xcafef00ddeadbeaf, 0xface864299792458},
			Severity:   0,
			Body:       "helo",
			Name:       "myname",
			NameSpace:  "nsps",
			InstanceId: Uuid{0xface864299792458, 0xcafef00ddeadbeaf},
			Version:    "1.0.1",
		},
	)
	if nil != e {
		log.Fatal(e)
	}
	log.Printf("converted: %s\n", converted)
}
