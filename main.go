package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/lusingander/geezer/geezer"
)

var (
	indentWidth = flag.Int("n", 2, "indent width")
	withSpace   = flag.String("s", "=", "characters with spaces before and after")
)

func run() error {
	flag.Parse()
	if *indentWidth < 0 {
		return fmt.Errorf("invalid indent width: %v", *indentWidth)
	}
	return geezer.Exec(os.Stdin, os.Stdout, *indentWidth, []rune(*withSpace))
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
