package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	APLHA   string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	NUMERIC string = "0123456789"
	SYMBOLS string = "~!@#$%^&*"
	ALL     string = NUMERIC + APLHA + SYMBOLS
)

func Generate(length int) []byte {
	rand.Seed(time.Now().UTC().UnixNano())
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] += ALL[rand.Intn(len(ALL))]
	}
	return result
}

func PrintGenerated(p []byte) {
	spacer := strings.Repeat("=", 24)

	fmt.Println(spacer)
	fmt.Println("Here's your password: 👇")
	fmt.Printf("%v\n", string(p))
	fmt.Println(spacer)
}
