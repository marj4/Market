package tests

import (
	"Market/backend/server"
	"testing"
)

func TestGenerateCodeForEmail(t *testing.T) {
	code, err := server.GenerateCodeForEmail()

	if err != nil {
		t.Fatalf("An error ocured,%v", err)
	}

	if len(code) == 0 || len(code) != 6 {
		t.Fatalf("Length of code should be 6, but got %v", len(code))
	}
}
