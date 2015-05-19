package main

import (
	"fmt"
	"os"
	"strconv"

	"./genpass"
)

func main() {

	var strong bool = false
	var length int = 14

	if len(os.Args) < 2 {
		fmt.Println("Usage: pass_gen <length> [-s]\t-s - optional, strong password.")
		return
	}

	if len(os.Args) == 3 && os.Args[2] == "-s" {
		strong = true
	}

	length, _ = strconv.Atoi(os.Args[1])
	password := genpass.Generate(length, strong)
	fmt.Println(password)
}