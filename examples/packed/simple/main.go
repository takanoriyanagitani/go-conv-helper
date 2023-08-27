package main

import (
	"context"
	"io"
	"log"

	"fmt"
	"strings"

	ch "github.com/takanoriyanagitani/go-conv-helper"
	pc "github.com/takanoriyanagitani/go-conv-helper/packed"
)

type input struct{ raw string }
type inputM struct{ packed [32]input }
type inputL struct{ nested []inputM }
type outM struct{ packed [32]output }
type outL struct{ nested []outM }

type output struct {
	lower string
	upper string
}

var conv ch.Converter[input, output] = func(_ context.Context, i input) (output, error) {
	return output{
		lower: strings.ToLower(i.raw),
		upper: strings.ToUpper(i.raw),
	}, nil
}

var convM ch.Converter[inputM, outM] = func(ctx context.Context, i inputM) (outM, error) {
	var out outM
	for j, input := range i.packed {
		o, e := conv(ctx, input)
		if nil != e {
			return out, e
		}
		out.packed[j] = o
	}
	return out, nil
}

var o2p pc.Output2Packed[outM, outL] = func(packed outL, converted outM) (outL, error) {
	return outL{nested: append(packed.nested, converted)}, nil
}

func dummyInputGen(i int) input {
	return input{raw: fmt.Sprintf("dummy Input %v", i)}
}

func dummyInputGenM(j int) inputM {
	var packed [32]input
	for i := range packed {
		packed[i] = dummyInputGen(i*j + i + j)
	}
	return inputM{packed}
}

func dummyInputGenL() inputL {
	var nested []inputM
	for j := 0; j < 8; j++ {
		nested = append(nested, dummyInputGenM(j))
	}
	return inputL{nested}
}

func inputSourceNew(i inputL) pc.InputSource[inputM] {
	var nested []inputM = i.nested
	var s []inputM = nested[:]
	return func() (inputM, error) {
		var sz int = len(s)
		if 0 == sz {
			return inputM{}, io.EOF
		}

		var ix int = sz - 1
		var last inputM = s[ix]
		s = s[:ix]
		return last, nil
	}
}

var i2p pc.Input2Packed[inputM, inputL] = func(output inputL, i inputM) (inputL, error) {
	return inputL{nested: append(output.nested, i)}, nil
}

var dummyInputSource pc.InputSource[inputM] = inputSourceNew(dummyInputGenL())

var unpackInput pc.UnpackPackedInput[inputL, inputM] = func(packed inputL) ([]inputM, error) {
	return packed.nested, nil
}

func main() {
	packed, e := i2p.PackAll(dummyInputSource, inputL{})
	if nil != e {
		log.Fatal(e)
	}

	var bld pc.ConvertPackedBuilder[inputL, inputM, outM, outL] = pc.
		ConvertPackedBuilder[inputL, inputM, outM, outL]{
		UnpackPackedInput: unpackInput,
		Converter:         convM,
		Output2Packed:     o2p,
	}
	var cp pc.ConvertPacked[inputL, inputM, outM, outL] = bld.Build()
	var ctx context.Context = context.Background()
	opack, e := cp(ctx, packed)
	if nil != e {
		log.Fatal(e)
	}
	log.Printf("opack: %v\n", opack)
}
