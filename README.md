Requirements
------------

Whatever LLVM, binaryen and wabt require; mostly cmake.


Setup
-----

- git submodule update --init --recursive
- make


Usage
-----

- bin/compile -o unit1.bc unit1.cpp
- bin/compile -o unit2.bc unit2.cpp
- bin/link -o program.wasm unit1.bc unit2.bc

