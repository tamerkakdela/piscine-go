package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	if nbr%2 == 1 {
		return false
	} else {
		return true
	}
}

func main() {
	Arg := os.Args[1:]
	if isEven(len(Arg)) {
		printStr(EvenMsg())
	} else {
		printStr(OddMsg())
	}
}

func EvenMsg() string {
	evenmsg := "I have an even number of arguments"
	return evenmsg
}

func OddMsg() string {
	oddmsg := "I have an odd number of arguments"
	return oddmsg
}
