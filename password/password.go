package password

import (
	"bufio"
	"errors"
	"fmt"
	"io/fs"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/ashalfarhan/passgo/icons"
	"github.com/ashalfarhan/passgo/utils"
	"github.com/atotto/clipboard"
	"github.com/fatih/color"
)

func Generate(length int, ns bool, nn bool) (*Password, error) {
	if length < 8 || length > 25 {
		return nil, errors.New("length must be at least 8 and less than 25")
	}
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
	return &Password{Value: result}, nil
}

func (p *Password) GetResult(save bool) string {
	spacer := strings.Repeat("=", 24)
	res := []string{
		spacer,
		color.GreenString("%v Here's your password:", icons.Donut),
		color.CyanString("%v %v", icons.Lock, string(p.Value)),
		spacer,
	}

	if p.Copied {
		res = append(res[:3], res[2:]...)
		res[3] = color.MagentaString("%v Copied to the clipboard", icons.Clip)
	}

	if save {
		res = append(res[:4], res[3:]...)
		res[3] = color.BlueString("%v Saved to: %s as: %s", icons.Paper, p.Filename, p.Passname)
	}

	return strings.Join(res, "\n")
}

func (p *Password) Copy() error {
	var err error = nil

	if clipboard.Unsupported {
		color.Yellow(icons.Warning + "Copy to clipboard is not supported")
		p.Copied = false
	} else {
		err = clipboard.WriteAll(string(p.Value))
		if err != nil {
			p.Copied = false
		} else {
			p.Copied = true
		}
	}

	return err
}

func (p *Password) Save() error {
	sc := bufio.NewReader(os.Stdin)
	filename, err := utils.GetInput("Filename (default \"passgo-generated\"): ", sc)
	if err != nil {
		return err
	}
	passname, err := utils.GetInput("Password name: ", sc)
	if err != nil {
		return err
	}

	if filename == "" {
		filename = "passgo-generated"
	}

	for passname == "" {
		color.Yellow("%v  Please enter a password name\n", icons.Warning)
		passname, err = utils.GetInput("Password name: ", sc)
		if err != nil {
			return err
		}
	}

	p.Filename, p.Passname = filename, passname
	file, err := os.OpenFile(p.Filename, os.O_APPEND|os.O_CREATE|os.O_RDWR, fs.ModeAppend)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("%s=%s\n", p.Passname, string(p.Value)))
	return err
}
