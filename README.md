This repository points to specific versions of third-party tools needed for
building C/C++ programs for the current version of
[wag](https://github.com/tsavola/wag).


Requirements
------------

Whatever LLVM requires: at least cmake, ninja-build and clang.
Linux, Go and libcapstone are required for testing.


Setup
-----

- `git clone https://github.com/tsavola/wag-toolchain.git`
- `git submodule update --init --recursive`
- `make`


Usage
-----

- `out/bin/clang --target=wasm32-unknown-unknown -emit-llvm -o unit1.bc unit1.cpp`
- `out/bin/clang --target=wasm32-unknown-unknown -emit-llvm -o unit2.bc unit2.cpp`
- `out/bin/clang -o program.wasm unit1.bc unit2.bc`


Testing
-------

- `go get -t github.com/tsavola/wag`
- `make check`
