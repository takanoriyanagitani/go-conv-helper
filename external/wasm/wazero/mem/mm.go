package memio

import (
	"context"
	"errors"

	wz "github.com/tetratelabs/wazero"
	wa "github.com/tetratelabs/wazero/api"

	util "github.com/takanoriyanagitani/go-conv-helper/util"

	ch "github.com/takanoriyanagitani/go-conv-helper"
	_ "github.com/takanoriyanagitani/go-conv-helper/external/wasm/wazero/simple"
)

var ErrUnableToReadOutput error = errors.New("unable to read output")
var ErrUnableToSetInput error = errors.New("unable to set input")
var ErrFunctionMissing error = errors.New("function missing")
var ErrMemoryMissing error = errors.New("memory missing")

type MemConverter struct {
	outputOffset uint32
	inputOffset  uint32
	memory       wa.Memory
	converter    func(ctx context.Context, isz int32) (osz int32, e error)
}

func (b MemConverter) getOutput(byteCount uint32) (output []byte, e error) {
	output, ok := b.memory.Read(b.outputOffset, byteCount)
	switch ok {
	case true:
		return output, nil
	default:
		return nil, ErrUnableToReadOutput
	}
}

func (b MemConverter) setInput(input []byte) error {
	var wrote bool = b.memory.Write(b.inputOffset, input)
	return util.Select(
		ErrUnableToSetInput,
		nil,
		wrote,
	)
}

func (b MemConverter) convert(ctx context.Context, isz int32) (output []byte, e error) {
	osz, e := b.converter(ctx, isz)
	return util.Select(
		func() ([]byte, error) { return nil, e },
		func() ([]byte, error) { return b.getOutput(uint32(osz)) },
		nil == e,
	)()
}

func (b MemConverter) Convert(ctx context.Context, input []byte) (output []byte, e error) {
	e = b.setInput(input)
	return util.Select(
		func() ([]byte, error) { return nil, e },
		func() ([]byte, error) { return b.convert(ctx, int32(len(input))) },
		nil == e,
	)()
}

func (b MemConverter) AsConverter() ch.RawConverter { return b.Convert }

type ConverterBuilder struct {
	output    wa.Function
	input     wa.Function
	converter wa.Function
	memory    wa.Memory
}

func (b ConverterBuilder) outputOffset(ctx context.Context) (uint32, error) {
	return b.getOffset(ctx, b.output)
}

func (b ConverterBuilder) inputOffset(ctx context.Context) (uint32, error) {
	return b.getOffset(ctx, b.input)
}

func (b ConverterBuilder) getOffset(ctx context.Context, f wa.Function) (uint32, error) {
	returns, e1 := f.Call(ctx)
	ret, e2 := b.toOffset(returns)
	return ret, errors.Join(e1, e2)
}

func (b ConverterBuilder) toOffset(raw []uint64) (uint32, error) {
	off, e := util.Array[uint64](raw).First()
	return uint32(wa.DecodeI32(off)), e
}

func (b ConverterBuilder) toConverter() func(ctx context.Context, isz int32) (osz int32, e error) {
	return func(ctx context.Context, isz int32) (osz int32, e error) {
		returns, e1 := b.converter.Call(ctx, wa.EncodeI32(isz))
		usz, e2 := util.Array[uint64](returns).First()
		osz = wa.DecodeI32(usz)
		return osz, errors.Join(e1, e2)
	}
}

func (b ConverterBuilder) GetOffset(ctx context.Context) (i, o uint32, e error) {
	i, ei := b.inputOffset(ctx)
	o, eo := b.outputOffset(ctx)
	return i, o, errors.Join(ei, eo)
}

func (b ConverterBuilder) Build(ctx context.Context) (MemConverter, error) {
	i, o, e := b.GetOffset(ctx)
	return MemConverter{
		outputOffset: o,
		inputOffset:  i,
		memory:       b.memory,
		converter:    b.toConverter(),
	}, e
}

type Module struct {
	raw      wa.Module
	outFunc  string
	inFunc   string
	convFunc string
	memName  string
}

func (m Module) getFunc(name string) (wa.Function, error) {
	var f wa.Function = m.raw.ExportedFunction(name)
	var e error = util.Select(
		nil,
		ErrFunctionMissing,
		nil == f,
	)
	return f, e
}

func (m Module) getMem(name string) (wa.Memory, error) {
	var mem wa.Memory = m.raw.ExportedMemory(name)
	var e error = util.Select(
		nil,
		ErrMemoryMissing,
		nil == mem,
	)
	return mem, e
}

func (m Module) getFunctions() (o wa.Function, i wa.Function, conv wa.Function, e error) {
	o, eo := m.getFunc(m.outFunc)
	i, ei := m.getFunc(m.inFunc)
	conv, ec := m.getFunc(m.convFunc)
	return o, i, conv, errors.Join(eo, ei, ec)
}

func (m Module) Build() (ConverterBuilder, error) {
	o, i, conv, ef := m.getFunctions()
	mem, em := m.getMem(m.memName)
	return ConverterBuilder{
		output:    o,
		input:     i,
		converter: conv,
		memory:    mem,
	}, errors.Join(ef, em)
}

type ModuleBuilder struct {
	compiled wz.CompiledModule
	config   wz.ModuleConfig
	runtime  wz.Runtime
	outFunc  string
	inFunc   string
	convFunc string
	memName  string
}

func (b ModuleBuilder) buildModule(ctx context.Context) (wa.Module, error) {
	return b.runtime.InstantiateModule(ctx, b.compiled, b.config)
}

func (b ModuleBuilder) Build(ctx context.Context) (Module, error) {
	m, e := b.buildModule(ctx)
	return Module{
		raw:      m,
		outFunc:  b.outFunc,
		inFunc:   b.inFunc,
		convFunc: b.convFunc,
		memName:  b.memName,
	}, e
}

func (b ModuleBuilder) Close(ctx context.Context) error { return b.runtime.Close(ctx) }

type RuntimeBuilder struct {
	rcfg     wz.RuntimeConfig
	mcfg     wz.ModuleConfig
	outFunc  string
	inFunc   string
	convFunc string
	memName  string
}

func (r RuntimeBuilder) buildRuntime(ctx context.Context) wz.Runtime {
	return wz.NewRuntimeWithConfig(ctx, r.rcfg)
}

func (r RuntimeBuilder) compile(
	ctx context.Context,
	wasm []byte,
	rtime wz.Runtime,
) (wz.CompiledModule, error) {
	return rtime.CompileModule(ctx, wasm)
}

func (r RuntimeBuilder) Build(ctx context.Context, wasm []byte) (ModuleBuilder, error) {
	var rt wz.Runtime = r.buildRuntime(ctx)
	compiled, e := r.compile(ctx, wasm, rt)
	return ModuleBuilder{
		compiled: compiled,
		config:   r.mcfg,
		runtime:  rt,
		outFunc:  r.outFunc,
		inFunc:   r.inFunc,
		convFunc: r.convFunc,
		memName:  r.memName,
	}, e
}

func (r RuntimeBuilder) WithRuntimeConfig(cfg wz.RuntimeConfig) RuntimeBuilder {
	r.rcfg = cfg
	return r
}

func (r RuntimeBuilder) WithModuleConfig(cfg wz.ModuleConfig) RuntimeBuilder {
	r.mcfg = cfg
	return r
}

func (r RuntimeBuilder) WithOutputFunctionName(name string) RuntimeBuilder {
	r.outFunc = name
	return r
}

func (r RuntimeBuilder) WithInputFunctionName(name string) RuntimeBuilder {
	r.inFunc = name
	return r
}

func (r RuntimeBuilder) WithConverterName(name string) RuntimeBuilder {
	r.convFunc = name
	return r
}

func (r RuntimeBuilder) WithMemName(name string) RuntimeBuilder {
	r.memName = name
	return r
}

var EmptyRuntimeBuilder RuntimeBuilder
var DefaultRuntimeBuilder RuntimeBuilder = EmptyRuntimeBuilder.
	WithRuntimeConfig(wz.NewRuntimeConfig()).
	WithModuleConfig(wz.NewModuleConfig()).
	WithOutputFunctionName("output_address").
	WithInputFunctionName("input_address").
	WithConverterName("converter").
	WithMemName("memory")
