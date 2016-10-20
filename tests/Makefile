hello: hello/hello.wasm
	go run test.go

%.bc: %.cpp
	../bin/compile -o $@ $*.cpp

%.wasm: %.bc
	../bin/link -o $@ $*.bc

.PHONY: hello
