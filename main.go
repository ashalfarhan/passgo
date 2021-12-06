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
	if err != nil {
		color.Red("%v  Failed to generate password: %v\n", icons.Warning, err.Error())
		os.Exit(1)
		return
	}

	err = pass.Copy()
	if err != nil {
		color.Red("%v  Failed to write password to the clipboard: %v\n", icons.Warning, err.Error())
		os.Exit(1)
		return
	}

	if save {
		err = pass.Save()
		if err != nil {
			color.Red("%v  Failed to save your password: %v\n", icons.Warning, err.Error())
			os.Exit(1)
			return
		}
	}

	res := pass.GetResult(save)
	fmt.Println(res)
}
