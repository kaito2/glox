package main

import "os"

func main() {
	l := Lox{}
	l.Main(os.Args[1:])
}
