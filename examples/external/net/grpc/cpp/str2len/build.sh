#!/bin/sh

genproto(){
	protoc \
		-I ./proto/str2len/v1 \
		--cpp_out=. \
		./proto/str2len/v1/sl.proto
}

gengrpc(){
	protoc \
		-I ./proto/str2len/v1 \
		--grpc_out=. \
		--plugin=protoc-gen-grpc=`which grpc_cpp_plugin` \
		./proto/str2len/v1/sl.proto
}

gen(){
	genproto
	gengrpc
}

buildapp(){
	export PKG_CONFIG_PATH=~/.local/grpc.d/lib/pkgconfig
	
	g++ \
		-std=c++23 \
		-Wall \
		-Werror \
		-O3 \
		-o app \
		-I ~/.local/grpc.d/include \
		-L ~/.local/grpc.d/lib \
		main.cpp \
		*.cc \
		$(
			pkg-config \
			--libs \
			--static \
			protobuf \
			grpc++ \
			grpc
		) \
		-lgrpc++_reflection \
		-ldl

}

gen
buildapp
