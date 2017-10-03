FROM debian:stretch

COPY bin/compile /usr/local/bin/
COPY bin/link /usr/local/bin/
COPY binaryen-build/bin/s2wasm /usr/local/binaryen-build/bin/
COPY llvm-build/bin/clang-6.0 /usr/local/llvm-build/bin/clang
COPY llvm-build/bin/llc /usr/local/llvm-build/bin/
COPY llvm-build/bin/llvm-as /usr/local/llvm-build/bin/
COPY llvm-build/bin/llvm-link /usr/local/llvm-build/bin/
COPY wabt/out/wast2wasm /usr/local/wabt/out/

RUN apt-get update && \
    DEBIAN_FRONTEND=noninteractive apt-get -y --no-install-recommends install python && \
    apt-get clean
