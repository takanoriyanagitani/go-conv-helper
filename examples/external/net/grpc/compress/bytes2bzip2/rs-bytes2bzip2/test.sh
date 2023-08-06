#!/bin/sh

grpcurl \
    -plaintext \
    -d '{"level":"LEVEL_NUM_FAST", "input":"aHcK"}' \
    -import-path $PWD \
    -proto b2b.proto \
    localhost:50051 \
    bytes2bz.v1.CompressService/Compress \
    | jq \
        --raw-output \
        .compressed \
    | base64 --decode \
    | bzcat
