#!/bin/sh

app_build(){
	g++ \
		-Wall \
		-Werror \
		-std=c++23 \
		-mavx2 \
		-O3 \
		-o app \
		./main.cpp \
		-lcbor
}

app_build
