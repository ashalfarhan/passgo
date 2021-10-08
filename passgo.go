package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func randInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}

func generate(length int) {
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = byte(randInt(33, 126))
	}
	fmt.Println(strings.Repeat("=", 48))
	fmt.Println("\tHere's your password: 👇")
	fmt.Printf("\t%v\n",string(result))
	fmt.Println(strings.Repeat("=", 48))
}

func printHelp() {
	fmt.Println(strings.Repeat("=", 48))
	fmt.Printf("[Passgo]\n\tUsage:\n")
	fmt.Println("\t (--l|--length)=<number> Enter specific number (default: 8)")
	fmt.Println("\t (--h|--help)            Print this guide")
	fmt.Println(strings.Repeat("=", 48))
}

func parseOption(a string) (opt, val string) {
	prefix := "--"
	args := strings.Split(strings.TrimPrefix(a, prefix), "=")
	opt = args[0]
	if len(args) == 2 {
		val = args[1]
	} else {
		val = ""
	}
	return
}

func main() {
	length := 8
	if len(os.Args) == 1 {
		generate(length)
		return
	}
	if len(os.Args) > 2 {
		fmt.Println("Too much flag")
		return
	}
	args := os.Args[1]
	opt, val := parseOption(args)
	switch opt {
	case "h", "help":
		printHelp()
	case "l", "length":
		lStr, err := strconv.ParseInt(val, 0, 32)
		if err != nil {
			fmt.Println("Invalid length flag input, we decide to generate default (8)")
			generate(length)
			return
		}
		length = int(lStr)
		generate(length)
	default:
		fmt.Printf("Invalid flag\n\t Try with --h flag to display help\n")
	}

}
