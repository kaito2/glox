package lox

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Lox struct {
	HasError bool
}

func (l *Lox) Main(args []string) {
	if len(args) > 1 {
		log.Fatal("usage: glox [script]")
	} else if len(args) == 1 {
		l.RunFile(args[0])
	} else {
		l.RunPrompt()
	}
}

func (l *Lox) RunFile(filePath string) error {
	b, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to ReadFile: %v", err)
	}
	if err := l.Run(b); err != nil {
		return fmt.Errorf("failed to Run: %v", err)
	}
	if l.HasError {
		os.Exit(65)
	}
	return nil
}

func (l *Lox) Run(b []byte) error {
	fmt.Println(string(b))
	return nil
}

func (l *Lox) Error(line int, message string) {
	l.Report(line, "", message)
}

func (l *Lox) Report(line int, where string, message string) {
	fmt.Printf("[line %d] Error%s: %s", line, where, message)
	l.HasError = true
}

func (l *Lox) RunPrompt() error {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("> ")
	for scanner.Scan() {
		l.Run(scanner.Bytes())
		l.HasError = false
		fmt.Print("> ")
	}

	return nil
}
