package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	ExitCodeOK int = iota
	ExitCodeError
)

var (
	Version  string
	Revision string
)

var app = App{}

func init() {
	version := flag.Bool("v", false, "prints current zoku version")
	help := flag.Bool("h", false, "prints usage.")
	flag.Parse()

	if *version {
		fmt.Println("Version: " + Version)
		fmt.Println("Revision: " + Revision)
		os.Exit(ExitCodeOK)
	}

	if *help {
		app.Usage()
		os.Exit(ExitCodeOK)
	}
}

func main() {
	os.Exit(app.Run(flag.Args()))
}
