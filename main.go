package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gkiryaziev/go-password-generator/generator"
)

func main() {

	var strong bool = false
	var length int = 14

	if len(os.Args) < 2 {
		fmt.Println("Usage: password-generator <length> [-s]\n\t-s - optional, strong password.")
		return
	}

	if len(os.Args) == 3 && os.Args[2] == "-s" {
		strong = true
	}

	length, _ = strconv.Atoi(os.Args[1])
	password := generator.Generate(length, strong)
	fmt.Println(password)
}
