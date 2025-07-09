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
	flag.BoolVar(&save, "save", false, "Whether save to a file or not")
	flag.BoolVar(&noSym, "no-symbols", false, "Exclude symbols")
	flag.BoolVar(&noNum, "no-numbers", false, "Exclude numbers")
	flag.Parse()
}

func main() {
	pass, err := password.Generate(length, noSym, noNum)
	if err != nil {
		PrintErr("Failed to generate password", err)
		return
	}

	if err = pass.Copy(); err != nil {
		PrintErr("Failed to write password to the clipboard", err)
	}

	if save {
		if err = pass.Save(); err != nil {
			PrintErr("Failed to save your password", err)
		}
	}

	fmt.Println(pass.GetResult())
}

func PrintErr(msg string, err error) {
	fmt.Fprint(os.Stderr, color.HiRedString("%v  %s: %v\n", icons.Warning, msg, err.Error()))
}
