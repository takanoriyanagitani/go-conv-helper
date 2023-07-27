package wzsimple

import (
	"context"
	"errors"

	wz "github.com/tetratelabs/wazero"
	wa "github.com/tetratelabs/wazero/api"

	wc "github.com/takanoriyanagitani/go-conv-helper/external/wasm"
	util "github.com/takanoriyanagitani/go-conv-helper/util"
)

var ErrInvalidReturn error = errors.New("invalid return")
var ErrFunctionMissing error = errors.New("function missing")

type SimpleFuncUU wc.SimpleFunc[uint64, uint64]

func ConvertFunc[I, O wc.SimpleType](
	encoder func(input I) (encoded uint64),
	decoder func(encoded uint64) (output O),
	original SimpleFuncUU,
) wc.SimpleFunc[I, O] {
	return func(ctx context.Context, input I) (output O, e error) {
		var encoded uint64 = encoder(input)
		raw, e := original(ctx, encoded)
		var decoded O = decoder(raw)
		return decoded, e
	}
}

type FuncConverter[I, O wc.SimpleType] func(SimpleFuncUU) wc.SimpleFunc[I, O]

var ConverterDD FuncConverter[float64, float64] = util.CurryIII(
	ConvertFunc[float64, float64],
)(wa.EncodeF64)(wa.DecodeF64)

var ConverterFF FuncConverter[float32, float32] = util.CurryIII(
	ConvertFunc[float32, float32],
)(wa.EncodeF32)(wa.DecodeF32)

var ConverterII FuncConverter[int32, int32] = util.CurryIII(
	ConvertFunc[int32, int32],
)(wa.EncodeI32)(wa.DecodeI32)

var ConverterUU FuncConverter[uint32, uint32] = util.CurryIII(
	ConvertFunc[uint32, uint32],
)(wa.EncodeU32)(wa.DecodeU32)

type Function struct{ internal wa.Function }

func GetOne[T any](returns []T) (t T, e error) {
	switch len(returns) {
	case 1:
		return returns[0], nil
	default:
		return t, ErrInvalidReturn
	}
}

func (f Function) CallUU(ctx context.Context, input uint64) (output uint64, e error) {
	ret, e := f.internal.Call(ctx, input)
	u, e2 := GetOne(ret)
	return u, errors.Join(e, e2)
}

func (f Function) ToSimpleFunc() SimpleFuncUU { return f.CallUU }

type Module struct {
	Name string
	wa.Module
}

func (b Module) Build() (f Function, e error) {
	var internal wa.Function = b.Module.ExportedFunction(b.Name)
	switch internal {
	case nil:
		return f, ErrFunctionMissing
	default:
		return Function{internal}, nil
	}
}

func (b Module) BuildSimpleFunc() (SimpleFuncUU, error) {
	f, e := b.Build()
	return util.Select(
		func() (SimpleFuncUU, error) { return nil, e },
		func() (SimpleFuncUU, error) { return f.ToSimpleFunc(), nil },
		nil == e,
	)()
}

type ModuleBuilder struct {
	wz.Runtime
	wz.CompiledModule
	wz.ModuleConfig
	FunctionName string
}

func (b ModuleBuilder) Build(ctx context.Context) (Module, error) {
	m, e := b.Runtime.InstantiateModule(
		ctx,
		b.CompiledModule,
		b.ModuleConfig,
	)
	return Module{
		Name:   b.FunctionName,
		Module: m,
	}, e
}

func (b ModuleBuilder) NewSimpleFuncBuilder() func(context.Context) (SimpleFuncUU, error) {
	return util.ComposeErr(
		b.Build,
		func(m Module) (SimpleFuncUU, error) { return m.BuildSimpleFunc() },
	)
}

type Runtime struct {
	wz.Runtime
	wz.ModuleConfig
	FunctionName string
}

func (c Runtime) Compile(ctx context.Context, wasm []byte) (ModuleBuilder, error) {
	cm, e := c.Runtime.CompileModule(ctx, wasm)
	return ModuleBuilder{
		Runtime:        c.Runtime,
		CompiledModule: cm,
		ModuleConfig:   c.ModuleConfig,
		FunctionName:   c.FunctionName,
	}, e
}

func (c Runtime) NewSimpleFuncBuilder() func(ctx context.Context, wasm []byte) (SimpleFuncUU, error) {
	return util.ComposeCtx(
		c.Compile,
		func(ctx context.Context, b ModuleBuilder) (SimpleFuncUU, error) {
			return b.NewSimpleFuncBuilder()(ctx)
		},
	)
}

type RuntimeBuilder struct {
	wz.ModuleConfig
	wz.RuntimeConfig
	FunctionName string
}

func (r RuntimeBuilder) Build(ctx context.Context) Runtime {
	var runtime wz.Runtime = wz.NewRuntimeWithConfig(ctx, r.RuntimeConfig)
	return Runtime{
		Runtime:      runtime,
		ModuleConfig: r.ModuleConfig,
		FunctionName: r.FunctionName,
	}
}

func (r RuntimeBuilder) NewSimpleFuncBuilder() wc.SimpleFuncBuilder[uint64, uint64] {
	return func(ctx context.Context, wasm []byte) (wc.SimpleFunc[uint64, uint64], error) {
		var rtime Runtime = r.Build(ctx)
		s, e := rtime.NewSimpleFuncBuilder()(ctx, wasm)
		return wc.SimpleFunc[uint64, uint64](s), e
	}
}

func (r RuntimeBuilder) WithModuleConfig(cfg wz.ModuleConfig) RuntimeBuilder {
	return RuntimeBuilder{
		ModuleConfig:  cfg,
		RuntimeConfig: r.RuntimeConfig,
		FunctionName:  r.FunctionName,
	}
}

func (r RuntimeBuilder) WithRuntimeConfig(cfg wz.RuntimeConfig) RuntimeBuilder {
	return RuntimeBuilder{
		ModuleConfig:  r.ModuleConfig,
		RuntimeConfig: cfg,
		FunctionName:  r.FunctionName,
	}
}

func (r RuntimeBuilder) WithFunctionName(name string) RuntimeBuilder {
	return RuntimeBuilder{
		ModuleConfig:  r.ModuleConfig,
		RuntimeConfig: r.RuntimeConfig,
		FunctionName:  name,
	}
}

var EmptyRuntimeBuilder RuntimeBuilder

var DefaultRuntimeBuilder RuntimeBuilder = EmptyRuntimeBuilder.
	WithModuleConfig(wz.NewModuleConfig()).
	WithRuntimeConfig(wz.NewRuntimeConfig()).
	WithFunctionName("")
