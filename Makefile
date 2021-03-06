J	:= $(shell getconf _NPROCESSORS_ONLN || echo 1)

CMAKE	:= cmake
NINJA	:= ninja

build:
	[ -L llvm/tools/clang ] || ln -s ../../clang llvm/tools/clang
	[ -L llvm/tools/lld ] || ln -s ../../lld llvm/tools/lld
	mkdir -p out
	cd out && $(CMAKE) -DCMAKE_BUILD_TYPE=MinSizeRel -DCMAKE_CROSSCOMPILING=True -DLLVM_DEFAULT_TARGET_TRIPLE=wasm32-unknown-unknown -DLLVM_TARGET_ARCH=wasm32 -G Ninja ../llvm
	$(NINJA) -j$(J) -C out bin/clang-8 bin/lld bin/llvm-as bin/llvm-link
	ln -sf clang-8 out/bin/clang
	- strip out/bin/*

check:
	$(MAKE) -C tests

clean:
	rm -rf out

.PHONY: build check clean
