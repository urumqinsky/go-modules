package main

import (
	"bufio"
	"os"

	"github.com/urumqinsky/go-modules/module00/ex00/internal/app/anscombe"
	"github.com/urumqinsky/go-modules/module00/ex00/internal/app/flag"
)

func main() {
	flags := flag.NewFlags()
	flags.Parse()

	data := anscombe.NewAnscombe()
	data.Scan(bufio.NewScanner(os.Stdin))
	data.Print(flags)
}
