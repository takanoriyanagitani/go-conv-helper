import { readFile } from "node:fs/promises";

const main = async () => {
	const wasm = await readFile(
		"./rs_memcopy.wasm",
		{
			encoding: null,
			flag: "r",
		},
	);
	const { instance, module } = await WebAssembly.instantiate(
		wasm,
		{},
	);
	const { exports } = instance;
	const {
		memory,
		input_address,
		output_address,
		converter,
	} = exports;
	const iaddr = input_address();
	const oaddr = output_address();
	const { buffer } = memory;
	const view = new Uint16Array(buffer, iaddr, 32768);
	view[0] = 3776;
	const osz = converter(4);
	const copied = new Uint16Array(buffer, oaddr, 32768);
	const out = copied[0];
	return {
		memory,
		buffer,
		iaddr,
		oaddr,
		osz,
		module,
		out,
	};
};

Promise.resolve()
.then(main)
.then(console.info)
.catch(console.error)
