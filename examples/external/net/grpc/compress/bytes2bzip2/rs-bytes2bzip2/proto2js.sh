#!/bin/sh

./node_modules/protobufjs-cli/bin/pbjs \
	--target static-module \
	--out ./b2b.proto.js \
	./b2b.proto
