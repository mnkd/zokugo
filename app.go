package main

import (
	"fmt"
	"os"
)

type App struct{}

func (app App) Usage() {
	str := `Usage:
 zokugo [-h|-v] [string]

Examples:
 $ zokugo あいうえお
 亜慰宇餌悪

 $ zokugo よろしく
 世露死苦
`
	fmt.Fprintln(os.Stderr, str)
}

func (app App) Run(args []string) int {
	if len(args) < 1 {
		app.Usage()
		return ExitCodeError
	}
	fmt.Fprintln(os.Stdout, Convert(args[0]))
	return ExitCodeOK
}
