package main

import (
	"context"
	"log"
	"os"

	ch "github.com/takanoriyanagitani/go-conv-helper"
	wm "github.com/takanoriyanagitani/go-conv-helper/external/wasm/wazero/mem"
)

func must[T any](t T, e error) T {
	switch e {
	case nil:
		return t
	default:
		panic(e)
	}
}

var rbld wm.RuntimeBuilder = wm.DefaultRuntimeBuilder
var rootCtx context.Context = context.Background()
var wasmBytes []byte = must(os.ReadFile("./rs_memcopy.wasm"))

func main() {
	var mbld wm.ModuleBuilder = must(rbld.Build(rootCtx, wasmBytes))
	defer mbld.Close(rootCtx)
	var mdl wm.Module = must(mbld.Build(rootCtx))
	var cbld wm.ConverterBuilder = must(mdl.Build())
	var mcv wm.MemConverter = must(cbld.Build(rootCtx))
	var cvt ch.RawConverter = mcv.AsConverter()
	var input []byte = []byte("hw")
	var output []byte = must(cvt(rootCtx, input))
	log.Printf("in: %v\n", input)
	log.Printf("out: %v\n", output)
}
