package main

import (
	"strings"
	"testing"

	"github.com/ashalfarhan/passgo/password"
)

func TestGeneratorValidation(t *testing.T) {
	t.Run("should return an error if length is greater than 25", func(t *testing.T) {
		if _, err := password.Generate(20000000, false, false); err == nil {
			t.Fatalf("\nexpected error, but got: %v", err)
		}
	})

	t.Run("should return an error if length is less than 8", func(t *testing.T) {
		if  _, err := password.Generate(-200, false, false); err == nil {
			t.Fatalf("\nexpected error, but got: %v", err)
		}
	})

	t.Run("should be success generate given length", func(t *testing.T) {
		ex := 12
		if p, _ := password.Generate(ex, false, false); ex != len(p.Value) {
			t.Fatalf("\nexpected result's length to be %v, but got: %v", ex, len(p.Value))
		}
	})

}

func TestGenerateInSequence(t *testing.T) {
	seq := 4

	t.Run("should return no symbol, if no-symbol is true", func(t *testing.T) {
		for i := 0; i < seq; i++ {
			p, _ := password.Generate(8, true, false)
			for _, j := range string(p.Value) {
				if strings.ContainsRune(password.SYMBOLS, j) {
					t.Fatalf("\nexpected generated is not contain symbol, but got: %v", string(j))
					return
				}
			}
		}
	})

	t.Run("should return no number, if no-number is true", func(t *testing.T) {
		for i := 0; i < seq; i++ {
			p, _ := password.Generate(8, false, true)
			for _, j := range string(p.Value) {
				if strings.ContainsRune(password.NUMERIC, j) {
					t.Fatalf("\nexpected generated is not contain number, but got: %v", string(j))
					return
				}
			}
		}
	})

	t.Run("should return no number and no symbol, if no-number and no-symbol is true", func(t *testing.T) {
		for i := 0; i < seq; i++ {
			p, _ := password.Generate(8, true, true)
			for _, j := range string(p.Value) {
				if strings.ContainsRune(password.NUMERIC, j) || strings.ContainsRune(password.SYMBOLS, j) {
					t.Fatalf("\nexpected generated is not contain number and symbol, but got: %v", string(j))
					return
				}
			}
		}
	})
}
