package packed_test

import (
	"io"
	"testing"

	"strings"

	pconv "github.com/takanoriyanagitani/go-conv-helper/packed"
)

func assertEqNew[T any](comp func(a, b T) (same bool)) func(a, b T) func(*testing.T) {
	return func(a, b T) func(*testing.T) {
		return func(t *testing.T) {
			t.Helper()
			var same bool = comp(a, b)
			if !same {
				t.Errorf("Unexpected value got.\n")
				t.Errorf("Expected: %v\n", b)
				t.Fatalf("Got:      %v\n", a)
			}
		}
	}
}

func assertEq[T comparable](a, b T) func(*testing.T) {
	comp := func(a, b T) (same bool) { return a == b }
	return assertEqNew(comp)(a, b)
}

func assertNil(e error) func(*testing.T) {
	return func(t *testing.T) {
		t.Helper()
		if nil != e {
			t.Fatalf("Must be nil: %v", e)
		}
	}
}

func assertErr(e error) func(*testing.T) {
	return func(t *testing.T) {
		t.Helper()
		if nil == e {
			t.Fatalf("must not nil")
		}
	}
}

func TestPacked(t *testing.T) {
	t.Parallel()

	t.Run("Input2Packed", func(t *testing.T) {
		t.Parallel()

		t.Run("PackAll", func(t *testing.T) {
			t.Parallel()

			t.Run("upper", func(t *testing.T) {
				t.Parallel()

				var out []string
				var inputs []string = []string{
					"Helo",
					"wrlD",
				}

				var i2p pconv.Input2Packed[string, []string] = func(
					o []string,
					input string,
				) ([]string, error) {
					var output string = strings.ToUpper(input)
					return append(o, output), nil
				}

				var isource pconv.InputSource[string] = func() (string, error) {
					var sz int = len(inputs)
					if 0 == sz {
						return "", io.EOF
					}

					var ix int = sz - 1
					var last string = inputs[ix]
					inputs = inputs[:ix]
					return last, nil
				}

				out, e := i2p.PackAll(isource, out)
				t.Run("no err", assertNil(e))
				t.Run("2 outputs", assertEq(len(out), 2))
				t.Run("1st item", assertEq(out[1], "HELO"))
				t.Run("2nd item", assertEq(out[0], "WRLD"))
			})
		})
	})
}
