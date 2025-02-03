package main

import (
	"fmt"
	"github.com/nanorand/nanorand"
)

func main() {
	code1 := GenerateCodeForEmail()
	code2 := code1
	fmt.Println(code1)
	fmt.Println(code2)
}

func GenerateCodeForEmail() string {
	code, _ := nanorand.Gen(6)
	return code
}
