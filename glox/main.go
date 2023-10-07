package main

import (
	"log"
	"os"

	"github.com/kaito2/glox/lox"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "glox",
		Usage: "Lox interpreter with Go.",
		Action: func(cCtx *cli.Context) error {
			l := lox.Lox{}
			l.RunPrompt()
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:  "run",
				Usage: "Run .lox file.",
				Action: func(cCtx *cli.Context) error {
					l := lox.Lox{}
					filePath := cCtx.Args().Get(0)
					l.RunFile(filePath)
					return nil
				},
			},
			{
				Name:  "repl",
				Usage: "Run in interactive mode.",
				Action: func(cCtx *cli.Context) error {
					l := lox.Lox{}
					l.RunPrompt()
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
