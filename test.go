package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func main() {
	pass1 := ""
	pass2 := ""

	fmt.Scan(&pass1)

	hash, err := bcrypt.GenerateFromPassword([]byte(pass1), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(hash))

	fmt.Scan(&pass2)

	if err := bcrypt.CompareHashAndPassword(hash, []byte(pass2)); err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(hash), string(hash2))

}
