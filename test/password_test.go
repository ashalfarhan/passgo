package main

import (
	"testing"

	"github.com/ashalfarhan/passgo/password"
)

func TestGenerator(t *testing.T) {
	t.Run("should return an error if length is greater than 25", func(t *testing.T) {
		_, err := password.Generate(1e2, false, false)
		if err == nil {
			t.Fatalf("\nexpected error, but got: %v", err)
		}
	})
	
	t.Run("should return an error if length is less than 8", func(t *testing.T) {
		_, err := password.Generate(-200, false, false)
		if err == nil {
			t.Fatalf("\nexpected error, but got: %v", err)
		}
	})
	
	t.Run("should be success generate given length", func(t *testing.T) {
		ex := 12
		p, _ := password.Generate(ex, false, false)
		if ex != len(p.Value) {
			t.Fatalf("\nexpected result's length to be %v, but got: %v", ex, len(p.Value))
		}
	})
}
