package main

import (
	"flag"
	"log"
	"os"

	"github.com/lusingander/geezer/geezer"
)

var (
	indentWidth = flag.Int("n", 2, "indent width")
)

func run() error {
	flag.Parse()
	return geezer.Exec(os.Stdin, os.Stdout, *indentWidth)
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
