import { readFile } from "node:fs/promises";
import output from "./output.js";

const { Output } = output.output.v1;

const main = async () => {
  const wasm = await readFile(
    "./rs_gz2json2proto.wasm",
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
  const json = JSON.stringify({
    device: {
      id: { id: "cafef00d-dead-beaf-face-864299792458" },
      name: { name: "helo" },
    },
    timestamp: { unixtime_seconds: 42 },
    current: { current: 2.99792458 },
    data: [
      { timestamp_nanos: 333, voltage: 0.599 },
      { timestamp_nanos: 634, voltage: 3.776 },
    ],
  });
  const encoder = new TextEncoder();
  const { buffer } = memory;
  const view = new Uint8Array(buffer, iadr);
  const { read, written } = encoder.encodeInto(json, view);
  const osz = converter(written);
  const serialized = new Uint8Array(memory.buffer, oadr);
  const parsed = Output.decode(serialized, osz);
  return parsed;
};

Promise.resolve()
  .then(main)
  .then(console.info)
  .catch(console.error);
