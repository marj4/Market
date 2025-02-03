package server

import (
	"testing"
)

func TestGenerateCodeForEmail(t *testing.T) {
	code, err := GenerateCodeForEmail()

	if err != nil {
		t.Fatalf("An error ocured,%v", err)
	}

	if len(code) == 0 || len(code) != 6 {
		t.Fatalf("Length of code should be 6, but got %v", len(code))
	}
}

func TestHash(t *testing.T) {
	_, hash, err := Hash("131313")

	//Check on error
	if err != nil {
		t.Fatalf("An error ocured,%v", err)
	}

	//Check hash length
	if len(hash) == 0 {
		t.Fatalf("The hash dont be empty, %v", hash)
	}

	//Check hash length(should be 60)
	if len(hash) != 60 {
		t.Fatalf("The len hash should be 60, no %d", len(hash))
	}

}

func TestSendCodeToEmail(t *testing.T) {
	//Simulation enter email and send code
	email := "legogo@gmail.com"
	code := "123456"

	if result := SendCodeToEmail(email, code); result != nil {
		t.Fatalf("An error ocured,%v", result)
	}

}
