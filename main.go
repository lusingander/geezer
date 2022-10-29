package main

import (
	"log"
	"os"

	"github.com/lusingander/geezer/geezer"
)

func run() error {
	return geezer.Exec(os.Stdin, os.Stdout, 2)
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
