#!/bin/sh

test_grpc(){
	testjson |
		grpcurl \
		    -plaintext \
		    -d @ \
		    -import-path $PWD \
		    -proto b2b.proto \
		    localhost:50051 \
		    bytes2bz.v1.CompressService/Compress \
		    | jq \
		        --raw-output \
		        .compressed \
		    | base64 --decode \
		    | bzcat
}

testjson(){
	echo '{}' |
		jq '{"level": "LEVEL_NUM_FAST", "input": "aHcK"}'
}

json2proto_with_header(){
	node test.mjs
}

json2proto2stdout(){
	testjson |
		json2proto_with_header
}

test_curl(){
	json2proto2stdout |
		curl \
			--header 'Content-Type: application/grpc-web' \
			--header 'Accept: application/grpc-web' \
			--data-binary @- \
			--fail \
			--show-error \
			--silent \
			http://localhost:50051/bytes2bz.v1.CompressService/Compress |
			tail --bytes=+8 |
			head --bytes=42 |
			bzcat
}

test_grpc
test_curl
