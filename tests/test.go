package main

import (
	"bufio"
	"log"
	"os"

	"github.com/tsavola/wag"
	"github.com/tsavola/wag/runner"
)

func main() {
	const (
		maxTextSize   = 65536
		maxRODataSize = 4096
		stackSize     = 4096
	)

	f, err := os.Open("hello/hello.wasm")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	wasm := bufio.NewReader(f)

	p, err := runner.NewProgram(maxTextSize, maxRODataSize)
	if err != nil {
		log.Fatal(err)
	}
	defer p.Close()

	var m wag.Module
	m.Load(wasm, runner.Env, p.Text, p.ROData, p.RODataAddr(), nil)
	p.Seal()
	p.SetData(m.Data())
	p.SetFunctionMap(m.FunctionMap())
	p.SetCallMap(m.CallMap())
	minMemorySize, maxMemorySize := m.MemoryLimits()

	r, err := p.NewRunner(minMemorySize, maxMemorySize, stackSize)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	result, err := r.Run(0, m.Signatures(), os.Stdout)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("result: 0x%x", uint32(result))

	if result != -1592745712 {
		os.Exit(1)
	}
}
