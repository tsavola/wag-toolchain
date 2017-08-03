package main

import (
	"bufio"
	"bytes"
	"flag"
	"log"
	"os"

	"github.com/tsavola/wag"
	"github.com/tsavola/wag/dewag"
	"github.com/tsavola/wag/runner"
)

func main() {
	const (
		maxTextSize   = 65536
		maxRODataSize = 4096
		stackSize     = 4096
	)

	filename := flag.String("file", "", "WebAssembly filename")
	saveText := flag.String("save-text", "", "save binary file")
	dumpText := flag.Bool("dump-text", false, "disassemble text section to stdout before executing test")
	expectResult := flag.Int("result", 0, "expected result value")
	flag.Parse()

	f, err := os.Open(*filename)
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

	m := wag.Module{
		MainSymbol: "main",
	}

	err = m.Load(wasm, runner.Env, bytes.NewBuffer(p.Text[:0]), p.ROData, p.RODataAddr(), nil)
	if err != nil {
		log.Fatal(err)
	}

	p.Seal()
	p.SetData(m.Data())
	p.SetFunctionMap(m.FunctionMap())
	p.SetCallMap(m.CallMap())
	minMemorySize, maxMemorySize := m.MemoryLimits()

	if *saveText != "" {
		f, err = os.Create(*saveText)
		defer f.Close()

		if _, err := f.Write(m.Text()); err != nil {
			log.Fatal(err)
		}
	}

	if *dumpText {
		dewag.PrintTo(os.Stdout, m.Text(), m.FunctionMap(), nil)
	}

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

	if result != int32(*expectResult) {
		os.Exit(1)
	}
}
