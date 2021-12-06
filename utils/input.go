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
