package manager

import (
	"context"

	ch "github.com/takanoriyanagitani/go-conv-helper"
	util "github.com/takanoriyanagitani/go-conv-helper/util"
)

type ManagerBuilder[I, O any] interface {
	Build() Manager[I, O]
}

type ManagerBuilderFn[I, O any] func() Manager[I, O]

func (b ManagerBuilderFn[I, O]) Build() Manager[I, O] { return b() }

type Manager[I, O any] interface {
	Build() ch.Converter[I, O]
	Error() error
}

type managerFn[I, O any] struct {
	builder func() ch.Converter[I, O]
	err     func() error
}

func (f managerFn[I, O]) Build() ch.Converter[I, O] { return f.builder() }
func (f managerFn[I, O]) Error() error              { return f.err() }

func ManagerNew[I, O any](
	builder func() ch.Converter[I, O],
	eb func() func() error,
) Manager[I, O] {
	return managerFn[I, O]{
		builder: builder,
		err:     eb(),
	}
}

//go:generate stringer -type=ErrorLevel
type ErrorLevel uint64

const (
	ErrorLevelUnspecified ErrorLevel = iota

	ErrorLevelOk // not an error

	ErrorLevelLo // normal error
	ErrorLevelHi // "high" level error: e.g, at least retry at higher-level required
)

type ManagedError struct {
	original error
	manager  error
}

func (m ManagedError) Error() string        { return m.original.Error() }
func (m ManagedError) OriginalError() error { return m.original }
func (m ManagedError) ManagerError() error  { return m.manager }

func (m ManagedError) bothError() bool { return nil != m.original && nil != m.manager }
func (m ManagedError) noError() bool   { return nil == m.original && nil == m.manager }
func (m ManagedError) Level() ErrorLevel {
	switch {
	case m.bothError():
		return ErrorLevelHi
	case nil != m.manager:
		return ErrorLevelHi
	case nil != m.original:
		return ErrorLevelLo
	case m.noError():
		return ErrorLevelOk
	default:
		return ErrorLevelUnspecified
	}
}

type ManagedConverter[I, O any] struct {
	Manager[I, O]
}

func (m ManagedConverter[I, O]) newConverter(original ch.Converter[I, O]) ch.Converter[I, O] {
	return func(ctx context.Context, input I) (output O, e error) {
		output, e = original(ctx, input)
		return util.Select(
			func() (O, error) {
				var manager error = m.Manager.Error()
				var me ManagedError = ManagedError{
					original: e,
					manager:  manager,
				}
				return output, me
			},
			func() (O, error) { return output, nil },
			nil == e,
		)()
	}
}

func (m ManagedConverter[I, O]) NewConverter() ch.Converter[I, O] {
	var original ch.Converter[I, O] = m.Manager.Build()
	return m.newConverter(original)
}
