package main

import (
	"context"
	"crypto/sha256"
	"crypto/sha512"
	"log"

	ch "github.com/takanoriyanagitani/go-conv-helper"
	util "github.com/takanoriyanagitani/go-conv-helper/util"

	ma "github.com/takanoriyanagitani/go-conv-helper/multi/async"
)

var data2sha256 ch.Converter[[]byte, [32]byte] = util.NoErr(sha256.Sum256)
var data2sha512 ch.Converter[[]byte, [64]byte] = util.NoErr(sha512.Sum512)

var data2s256v ch.Converter[[]byte, SumsUpd] = ch.ConverterAdd(
	data2sha256,
	util.NoErr(func(value [32]byte) SumsUpd { return V256{value} }),
)

var data2s512v ch.Converter[[]byte, SumsUpd] = ch.ConverterAdd(
	data2sha512,
	util.NoErr(func(value [64]byte) SumsUpd { return V512{value} }),
)

type Sums struct {
	s256 [32]byte
	s512 [64]byte
}

type V256 struct{ value [32]byte }

func (v V256) Update(s Sums) Sums {
	return Sums{
		s256: v.value,
		s512: s.s512,
	}
}
func (v V256) AsIf() SumsUpd { return v }

type V512 struct{ value [64]byte }

func (v V512) Update(s Sums) Sums {
	return Sums{
		s256: s.s256,
		s512: v.value,
	}
}
func (v V512) AsIf() SumsUpd { return v }

type SumsUpd interface{ Update(s Sums) Sums }

func reducer(state Sums, next ma.Either[SumsUpd]) (Sums, error) {
	var e error = next.Left()
	if nil != e {
		return state, e
	}
	var upd SumsUpd = next.Right()
	var updated Sums = upd.Update(state)
	return updated, nil
}

var mconv ch.Converter[[]byte, Sums] = ma.ConverterMultiNew(
	reducer,
	ma.ConvertToChanNew(data2s256v),
	ma.ConvertToChanNew(data2s512v),
)

func main() {
	var input []byte = []byte("helo")
	sums, e := mconv(context.Background(), input)
	if nil != e {
		log.Fatal(e)
	}
	log.Printf("256: %v\n", sums.s256)
	log.Printf("512: %v\n", sums.s512)
}
