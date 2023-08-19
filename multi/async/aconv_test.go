package aconv_test

import (
	"context"
	"testing"

	"strings"

	ch "github.com/takanoriyanagitani/go-conv-helper"
	aconv "github.com/takanoriyanagitani/go-conv-helper/multi/async"
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

func TestConv(t *testing.T) {
	t.Parallel()

	t.Run("ConvertToChan", func(t *testing.T) {
		t.Parallel()

		t.Run("lower", func(t *testing.T) {
			t.Parallel()

			var lower ch.Converter[string, string] = func(
				_ context.Context,
				input string,
			) (string, error) {
				return strings.ToLower(input), nil
			}

			var c2c aconv.ConvertToChan[string, string] = aconv.ConvertToChanNew(lower)
			var out chan aconv.Either[string] = make(chan aconv.Either[string], 1)
			c2c(context.Background(), "HeLo", out)
			var result aconv.Either[string] = <-out
			var err error = result.Left()
			t.Run("no err", assertNil(err))
			var right string = result.Right()
			t.Run("same", assertEq(right, "helo"))
		})
	})

	t.Run("ConverterMultiNew", func(t *testing.T) {
		t.Parallel()

		t.Run("lower/upper", func(t *testing.T) {
			t.Parallel()

			var lower ch.Converter[string, string] = func(
				_ context.Context,
				input string,
			) (string, error) {
				return strings.ToLower(input), nil
			}

			var upper ch.Converter[string, string] = func(
				_ context.Context,
				input string,
			) (string, error) {
				return strings.ToUpper(input), nil
			}

			var s2l aconv.ConvertToChan[string, string] = aconv.ConvertToChanNew(lower)
			var s2h aconv.ConvertToChan[string, string] = aconv.ConvertToChanNew(upper)

			reducer := func(state []string, next aconv.Either[string]) ([]string, error) {
				var err error = next.Left()
				if nil != err {
					return nil, err
				}
				var right string = next.Right()
				return append(state, right), nil
			}

			var conv ch.Converter[string, []string] = aconv.ConverterMultiNew(
				reducer,
				s2l,
				s2h,
			)

			out, e := conv(
				context.Background(),
				"helo, WRLD",
			)

			t.Run("no err", assertNil(e))
			t.Run("two strings", assertEq(len(out), 2))
			t.Run("upper check", assertEq(out[0], "HELO, WRLD"))
			t.Run("lower check", assertEq(out[1], "helo, wrld"))
		})
	})
}
