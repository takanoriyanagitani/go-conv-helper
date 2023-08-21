#!/bin/sh

protodir=proto/str2len/v1
protofile=sl.proto
addr=localhost:50051
unix=./sock.d/test.sock

tcp(){
	echo '{}' |
	jq -c '{
		target: {
			id: {
				hi: 3405705229,
				lo: 3735928495
			}
		}
	}' |
	grpcurl \
		-plaintext \
		-d @ \
		-import-path "${protodir}" \
		-proto "${protofile}" \
		"${addr}" \
		str2len.v1.GetCmdService/GetCmd
}

uds(){
	echo '{}' |
	jq -c '{
		target: {
			id: {
				hi: 3405705229,
				lo: 3735928495
			}
		}
	}' |
	grpcurl \
		-plaintext \
		-d @ \
		-import-path "${protodir}" \
		-proto "${protofile}" \
		-unix=true \
		-authority "dummy" \
		"${unix}" \
		str2len.v1.GetCmdService/GetCmd
}

uds
