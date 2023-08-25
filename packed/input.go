package packed

import (
	"io"
)

// InputSource returns an input if exists.
// error must be io.EOF when there's no more data
type InputSource[I any] func() (I, error)

type Input2Packed[I, P any] func(output P, input I) (P, error)

func (c Input2Packed[I, P]) PackAll(inputs InputSource[I], output P) (P, error) {
	for {
		input, e := inputs()
		if io.EOF == e {
			return output, nil
		}
		if nil != e {
			return output, e
		}

		output, e = c(output, input)
		if nil != e {
			return output, e
		}
	}
}
