package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("File name missing")
	}
	if len(args) > 1 {
		fmt.Println("Too many arguments")
	}
	if len(args) == 1 {
		fmt.Println("Almost there!!")
	}
}
