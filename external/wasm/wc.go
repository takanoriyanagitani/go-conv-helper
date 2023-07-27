package wc

import (
	"context"

	ch "github.com/takanoriyanagitani/go-conv-helper"
)

type ConverterBuilder func(ctx context.Context, wasmBytes []byte) (ch.RawConverter, error)

type MemoryConfig struct {
	PageMax uint32
}

type Config struct {
	MemoryConfig
}
