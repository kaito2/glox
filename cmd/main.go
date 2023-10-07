package main

import (
	"os"

	"github.com/kaito2/glox/lox"
)

func main() {
	l := lox.Lox{}
	l.Main(os.Args[1:])
}
