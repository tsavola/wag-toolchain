J	:= 1
CMAKE	:= cmake

all: llvm binaryen wabt

llvm:
	[ -L llvm/tools/clang ] || ln -s ../../clang llvm/tools/clang
	mkdir -p llvm-build
	cd llvm-build && $(CMAKE) -DCMAKE_CROSSCOMPILING=True -DLLVM_DEFAULT_TARGET_TRIPLE=wasm32-unknown-unknown -DLLVM_TARGET_ARCH=wasm32 -DLLVM_TARGETS_TO_BUILD= -DLLVM_EXPERIMENTAL_TARGETS_TO_BUILD=WebAssembly ../llvm
	$(MAKE) -j$(J) -C llvm-build

binaryen:
	mkdir -p binaryen-build
	cd binaryen-build && $(CMAKE) ../binaryen
	$(MAKE) -j$(J) -C binaryen-build

wabt:
	$(MAKE) -j$(J) -C wabt

check:
	$(MAKE) -C tests

.PHONY: all llvm binaryen wabt tests
