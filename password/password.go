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
		comb = SYMBOLS + ALPHA
	} else if ns {
		comb = NUMERIC + ALPHA
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

func (p *Password) GetResult() string {
	spacer := strings.Repeat("=", 24)
	res := []string{
		spacer,
		color.GreenString("%v Here's your password:", icons.Donut),
		color.CyanString("%v %v", icons.Lock, string(p.Value)),
		spacer,
	}

	if p.Copied {
		res = utils.InsertAt(
			res,
			len(res)-1,
			color.MagentaString("%v Copied to the clipboard", icons.Clip),
		)
	}

	if p.Saved {
		res = utils.InsertAt(
			res,
			len(res)-1,
			color.HiBlueString("%v Saved to: %s as: %s", icons.Paper, p.Filename, p.Passname),
		)
	}

	return strings.Join(res, "\n")
}

func (p *Password) Copy() error {

	if clipboard.Unsupported {
		color.Yellow(icons.Warning + "Copy to clipboard is not supported")
		return nil
	}

	if err := clipboard.WriteAll(string(p.Value)); err != nil {
		return err
	}

	p.Copied = true
	return nil
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
		if passname, err = utils.GetInput("Password name: ", sc); err != nil {
			return err
		}
	}

	p.Filename, p.Passname = filename, passname
	file, err := os.OpenFile(p.Filename, os.O_APPEND|os.O_CREATE|os.O_RDWR, fs.ModePerm)
	if err != nil {
		return err
	}

	defer file.Close()

	if _, err = file.WriteString(fmt.Sprintf("%s=%s\n", p.Passname, string(p.Value))); err != nil {
		return err
	}
	
	p.Saved = true
	return nil
}
