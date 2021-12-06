package utils

import (
	"bufio"
	"fmt"
	"strings"
)

func GetInput(q string, r *bufio.Reader) (string, error) {
	fmt.Print(q)
	res, err := r.ReadString('\n')
	return strings.TrimSpace(res), err
}

// This should be generic
func InsertAt(arr []string, i int, el string) []string {
	arr = append(arr[:i], arr[i-1:]...)
	arr[i] = el
	return arr
}
