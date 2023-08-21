#!/bin/sh

addr=localhost:2151

echo '{}' |
jq -c '{
  meta: {
    reply_id: {
      hi: 65535,
      lo: 51966,
    },
  },
  hsv: {
    hue: 1.0,
    saturation: 0.5,
    value: 0.25,
  },
}' |
grpcurl \
    -plaintext \
    -d @ \
    -import-path proto/color/hsv/v1 \
    -proto conv.proto \
    "${addr}" \
    color.hsv.v1.HsvCmdService/Convert
