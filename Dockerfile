FROM ubuntu:artful

COPY out/bin/clang-7.0 /usr/local/out/bin/clang
COPY out/bin/llvm-as /usr/local/out/bin/
COPY out/bin/llvm-link /usr/local/out/bin/
COPY out/bin/wasm-ld /usr/local/out/bin/
