package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:]
	for _, str := range arguments {
		for _, s := range str {
			z01.PrintRune(s)
		}
		z01.PrintRune('\n')
	}
}
