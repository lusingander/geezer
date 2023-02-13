package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/lusingander/geezer/geezer"
)

var (
	indentWidth    = flag.Int("n", 2, "indent width")
	withSpaceRunes = []rune{'='}
)

func run() error {
	flag.Parse()
	if *indentWidth < 0 {
		return fmt.Errorf("invalid indent width: %v", *indentWidth)
	}
	return geezer.Exec(os.Stdin, os.Stdout, *indentWidth, withSpaceRunes)
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
