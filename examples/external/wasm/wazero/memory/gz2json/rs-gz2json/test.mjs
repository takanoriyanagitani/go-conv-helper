import { readFile } from "node:fs/promises";
import { promisify } from "node:util";
import { gzip } from "node:zlib";
import { createGzip } from "node:zlib";

const main = async () => {
  const wasm = await readFile(
    "./rs_gz2json.wasm",
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
  const iadr = input_address();
  const oadr = output_address();
  const input = "hw";
  const compressed = await promisify(gzip)(input);
  const isz = compressed.length;
  const { buffer } = memory;
  const view = new Uint8Array(buffer, iadr, 65536);
  const copied = compressed.copy(view);
  const osz = converter(isz);
  const oview = new Uint8Array(memory.buffer, oadr, osz);
  const bytes2string = new TextDecoder();
  const decoded = bytes2string.decode(oview);
  return {
    input,
    decoded,
  };
};

Promise.resolve()
  .then(main)
  .then(console.info)
  .catch(console.error);
