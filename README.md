This repository points to specific versions of third-party tools needed for
building C/C++ programs for the current version of
[wag](https://github.com/tsavola/wag).


Requirements
------------

Whatever LLVM, binaryen and wabt require: at least cmake and clang.
Linux, Go and libcapstone are required for testing.


Setup
-----

- `git clone https://github.com/tsavola/wag-toolchain.git`
- `git submodule update --init --recursive`
- `make`


Usage
-----

- `bin/compile -o unit1.bc unit1.cpp`
- `bin/compile -o unit2.bc unit2.cpp`
- `bin/link -o program.wasm unit1.bc unit2.bc`


Testing
-------

- `go get -t github.com/tsavola/wag`
- `make check`
