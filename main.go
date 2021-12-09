package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ashalfarhan/passgo/icons"
	"github.com/ashalfarhan/passgo/password"
	"github.com/fatih/color"
)

var (
	save   bool
	length int
	noSym  bool
	noNum  bool
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s\n", "passgo")
		flag.PrintDefaults()
	}
	flag.IntVar(&length, "length", 8, "Generated password length")
	flag.BoolVar(&save, "save", false, "Whether save to file or not")
	flag.BoolVar(&noSym, "no-symbol", false, "Except symbol")
	flag.BoolVar(&noNum, "no-number", false, "Except number")
	flag.Parse()
}

func main() {
	pass, err := password.Generate(length, noSym, noNum)
	ErrorExit("Failed to generate password", err)

	ErrorExit("Failed to write password to the clipboard", pass.Copy())

	if save {
		ErrorExit("Failed to save your password", pass.Save())
	}

	fmt.Println(pass.GetResult())
}

func ErrorExit(msg string, err error) {
	if err != nil {
		fmt.Fprint(os.Stderr, color.HiRedString("%v  %s: %v\n", icons.Warning, msg, err.Error()))
		os.Exit(1)
	}
}
