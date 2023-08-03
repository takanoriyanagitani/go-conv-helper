package main

import (
	"context"
	"log"
	"os"

	wc "github.com/takanoriyanagitani/go-conv-helper/external/wasm"
	ws "github.com/takanoriyanagitani/go-conv-helper/external/wasm/wazero/simple"
)

func must[T any](t T, e error) T {
	switch e {
	case nil:
		return t
	default:
		panic(e)
	}
}

var rbld ws.RuntimeBuilder = ws.DefaultRuntimeBuilder.
	WithFunctionName("sinf64")

var wasmBytes []byte = must(os.ReadFile("./rs_sin.wasm"))

var rootCtx context.Context = context.Background()

func main() {
	var rtime ws.Runtime = rbld.Build(rootCtx)
	defer rtime.Close(rootCtx)
	var wasm2sfu func(ctx context.Context, wasm []byte) (ws.SimpleFuncUU, error) = rtime.
		NewSimpleFuncBuilder()
	var sfuu ws.SimpleFuncUU = must(wasm2sfu(rootCtx, wasmBytes))
	var s64d wc.SimpleFunc[float64, float64] = ws.ConverterDD(sfuu)

	log.Printf("sin 0.0: %v\n", must(s64d(rootCtx, 0.0)))
	log.Printf("sin 3.1: %v\n", must(s64d(rootCtx, 3.1)))
}
