#!/bin/sh

benchcnt=16777216

\time --verbose go \
	test \
	-bench \
	Benchmark \
	-benchmem \
	-benchtime ${benchcnt}x \
	./...
