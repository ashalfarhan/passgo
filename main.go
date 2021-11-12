package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	save   bool
	length int
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s\n", "passgo")
		flag.PrintDefaults()
	}
	flag.BoolVar(&save, "save", false, "Whether save to file or not")
	flag.IntVar(&length, "length", 8, "Generated password length")
	flag.Parse()
}

func main() {
	pass := Generate(length)
	if save {
		flag.PrintDefaults()
		fmt.Println("This still beta")
		os.Exit(1)
	}
	PrintGenerated(pass)
	os.Exit(0)
}
