GO	?= go
RUNTEST	:= $(GO) run test.go

check: hello/hello.wasm
	$(RUNTEST) -file=hello/hello.wasm -result=-1592745712

%.bc: %.c
	../out/bin/clang --target=wasm32 -emit-llvm -o $@ $*.c

%.bc: %.cpp
	../out/bin/clang --target=wasm32 -emit-llvm -o $@ $*.cpp

%.wasm: %.bc
	../out/bin/clang -o $@ $*.bc

.PHONY: check
