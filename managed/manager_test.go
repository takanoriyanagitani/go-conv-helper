package manager_test

import (
	"context"
	"testing"

	"errors"
	"fmt"

	ch "github.com/takanoriyanagitani/go-conv-helper"
	managed "github.com/takanoriyanagitani/go-conv-helper/managed"
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

func TestManager(t *testing.T) {
	t.Parallel()

	t.Run("ManagedConverter", func(t *testing.T) {
		t.Parallel()

		t.Run("ManagerNew", func(t *testing.T) {
			t.Parallel()

			var mngr managed.Manager[string, string] = managed.ManagerNew(
				func() ch.Converter[string, string] {
					var dummyStateLo uint8 = 3
					return func(_ context.Context, input string) (string, error) {
						switch dummyStateLo {
						case 0:
							return "", fmt.Errorf("normal error")
						default:
							dummyStateLo -= 1
							return input, nil
						}
					}
				},
				func() func() error {
					var dummyStateHi uint8 = 5
					return func() error {
						switch dummyStateHi {
						case 0:
							return fmt.Errorf("fatal error")
						default:
							dummyStateHi -= 1
							return nil
						}
					}
				},
			)
			var mcnv managed.ManagedConverter[string, string] = managed.ManagedConverter[string, string]{
				Manager: mngr,
			}

			var ctx context.Context = context.Background()

			for i := 0; i < 5; i++ {
				var neo ch.Converter[string, string] = mcnv.NewConverter()
				o1, e1 := neo(ctx, "helo")
				t.Run("no err 1", assertNil(e1))
				t.Run("same out 1", assertEq(o1, "helo"))

				o2, e2 := neo(ctx, "wrld")
				t.Run("no err 2", assertNil(e2))
				t.Run("same out 2", assertEq(o2, "wrld"))

				o3, e3 := neo(ctx, "hw")
				t.Run("no err 3", assertNil(e3))
				t.Run("same out 3", assertEq(o3, "hw"))

				_, e4 := neo(ctx, "hh")
				t.Run("must err 4", assertErr(e4))
				var em managed.ManagedError
				if errors.As(e4, &em) {
					switch em.Level() {
					case managed.ErrorLevelLo:
						continue
					default:
						t.Fatalf("Unexpected error level: %v\n", em.Level().String())
					}
					continue
				}
				t.Fatalf("Unexpected err: %v\n", e4)
			}

			for i := 0; i < 1; i++ {
				var neo ch.Converter[string, string] = mcnv.NewConverter()
				o1, e1 := neo(ctx, "helo")
				t.Run("no err 1", assertNil(e1))
				t.Run("same out 1", assertEq(o1, "helo"))

				o2, e2 := neo(ctx, "wrld")
				t.Run("no err 2", assertNil(e2))
				t.Run("same out 2", assertEq(o2, "wrld"))

				o3, e3 := neo(ctx, "hw")
				t.Run("no err 3", assertNil(e3))
				t.Run("same out 3", assertEq(o3, "hw"))

				_, e4 := neo(ctx, "hh")
				t.Run("must err 4", assertErr(e4))
				var em managed.ManagedError
				if errors.As(e4, &em) {
					switch em.Level() {
					case managed.ErrorLevelHi:
						continue
					default:
						t.Fatalf("Unexpected error level: %v\n", em.Level().String())
					}
					continue
				}
				t.Fatalf("Unexpected err: %v\n", e4)
			}
		})
	})
}
