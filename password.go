package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	ALPHA   string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	NUMERIC string = "0123456789"
	SYMBOLS string = "~!@#$%^&*"
	ALL     string = NUMERIC + ALPHA + SYMBOLS
)

func Generate(length int, ns bool, nn bool) *Password {
	var comb string
	if ns && nn {
		comb = ALPHA
	} else if nn {
		comb = ALPHA + SYMBOLS
	} else if ns {
		comb = ALPHA + NUMERIC
	} else {
		comb = ALL
	}
	rand.Seed(time.Now().UTC().UnixNano())
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = comb[rand.Intn(len(comb))]
	}
	return &Password{value: result}
}

type Password struct {
	value    []byte
	filename string
	passname string
}

func (p *Password) PrintGenerated(s bool) {
	spacer := strings.Repeat("=", 24)
	res := []string{spacer, "Here's your password: 👇", fmt.Sprintf("%v", string(p.value)), spacer}
	if s {
		res = append(res[:3], res[2:]...)
		res[3] = fmt.Sprintf("Saved to: %s as: %s", p.filename, p.passname)
	}
	fmt.Print(strings.Join(res, "\n"))
}

func (p *Password) Save() {
	sc := bufio.NewReader(os.Stdin)
	filename, _ := GetInput("Filename (default \"passgo-generated\"): ", sc)
	passname, _ := GetInput("Password name (default current date): ", sc)
	if filename == "" {
		filename = "passgo-generated"
	}
	p.filename = filename

	if passname == "" {
		passname = time.Now().UTC().String()
	}
	p.passname = passname

	file, err := os.OpenFile(p.filename, os.O_APPEND|os.O_CREATE, 0755)
	if err != nil {
		log.Fatalln(err)
	}

	defer file.Close()
	file.WriteString(fmt.Sprintf("%s=%s\n", p.passname, string(p.value)))
	p.PrintGenerated(true)
}

func GetInput(q string, r *bufio.Reader) (string, error) {
	fmt.Print(q)
	res, err := r.ReadString('\n')
	return strings.TrimSpace(res), err
}

// Get current dir
// home, _ := os.UserHomeDir()
// fmt.Println(filepath.Join(home, filename))
