#!/bin/sh

python3 -m grpc_tools.protoc \
  -I ./proto/color/hsv/v1 \
  --python_out=. \
  --pyi_out=. \
  --grpc_python_out=. \
  ./proto/color/hsv/v1/conv.proto
