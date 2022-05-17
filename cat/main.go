package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) >= 1 {
		for _, file := range args {
			data, err := os.ReadFile(file)
			if err != nil {
				printTheFile("ERROR: open " + file + ": no such file or directory")
				z01.PrintRune('\n')
				os.Exit(1)
			} else {
				printTheFile(string(data))
			}

		}
	} else {
		str, _ := ioutil.ReadAll(os.Stdin)
		printTheFile(string(str))
	}
}

func printTheFile(str string) {
	for _, i := range str {
		z01.PrintRune(i)
	}
}
