import { readFile, writeFile } from "node:fs/promises";

import b2b from "./b2b.proto.js";

const {
  Compression,
  CompressRequest,
  CompressResponse,
  CompressService,
} = b2b.bytes2bz.v1;

const l2l = (lstring) => {
  switch (lstring) {
    default:
      return Compression.Level.LEVEL_NUM_FAST;
  }
};

const main = async () => {
  const json = await readFile(
    "/dev/stdin",
    {
      encoding: "utf8",
      flag: "r",
    },
  );
  const parsed = JSON.parse(json);
  const {
    level,
    input,
  } = parsed;
  const rawInput = Buffer.from(input, "base64");
  const request = CompressRequest.create({
    level: l2l(level),
    input: rawInput,
  });
  const serialized = CompressRequest.encode(request).finish();
  const deserialized = CompressRequest.decode(serialized);
  const length = serialized.length;
  const flag = Buffer.allocUnsafe(1);
  flag[0] = 0;
  const contentLength = Buffer.allocUnsafe(4);
  contentLength.writeUInt32BE(length);
  const header = Buffer.concat([flag, contentLength]);
  const full = Buffer.concat([header, serialized]);
  return writeFile(
    "/dev/stdout",
    full,
    {
      encoding: null,
    },
  );
};

Promise.resolve()
  .then(main)
  .catch(console.error);
