package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	save   bool
	length int
	noSymbol bool
	noNum bool
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s\n", "passgo")
		flag.PrintDefaults()
	}
	flag.IntVar(&length, "length", 8, "Generated password length")
	flag.BoolVar(&save, "save", false, "Whether save to file or not")
	flag.BoolVar(&noSymbol, "no-symbol", false, "Except symbol")
	flag.BoolVar(&noNum, "no-number", false, "Except number")
	flag.Parse()
}

func main() {
	pass := Generate(length, noSymbol, noNum)
	if save {
		pass.Save()
		os.Exit(0)
		return
	}
	pass.PrintGenerated(false)
}
