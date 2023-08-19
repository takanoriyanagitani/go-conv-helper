package conv

import (
	util "github.com/takanoriyanagitani/go-conv-helper/util"
)

func ConverterAdd[I, O1, O2 any](c1 Converter[I, O1], c2 Converter[O1, O2]) Converter[I, O2] {
	return util.ComposeCtx(
		c1,
		c2,
	)
}
