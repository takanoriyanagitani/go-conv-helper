package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	ch "github.com/takanoriyanagitani/go-conv-helper"
	ms "github.com/takanoriyanagitani/go-conv-helper/multi/sync"
)

var str2hi ch.Converter[string, string] = func(_ context.Context, input string) (string, error) {
	return strings.ToUpper(input), nil
}

var str2lo ch.Converter[string, string] = func(_ context.Context, input string) (string, error) {
	return strings.ToLower(input), nil
}

type LoHi struct {
	lo string
	hi string
}

var finalizer ch.Converter[[]string, LoHi] = func(_ context.Context, input []string) (LoHi, error) {
	if 2 != len(input) {
		return LoHi{}, fmt.Errorf("Unexpected slice length: %v", len(input))
	}
	return LoHi{
		lo: input[0],
		hi: input[1],
	}, nil
}

var str2lh ch.Converter[string, LoHi] = ms.ConverterMultiNew(
	finalizer,
	str2lo,
	str2hi,
)

func main() {
	lh, e := str2lh(context.Background(), "Helo")
	if nil != e {
		log.Fatal(e)
	}
	log.Printf("lo: %s\n", lh.lo)
	log.Printf("hi: %s\n", lh.hi)
}
