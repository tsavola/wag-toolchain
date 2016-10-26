GO	?= go
RUNTEST	:= $(GO) run test.go

check: hello/hello.wasm
	$(RUNTEST) -file=hello/hello.wasm -result=-1592745712

%.bc: %.c
	../bin/compile -o $@ $*.c

%.bc: %.cpp
	../bin/compile -o $@ $*.cpp

%.wasm: %.bc
	../bin/link -o $@ $*.bc

.PHONY: check
