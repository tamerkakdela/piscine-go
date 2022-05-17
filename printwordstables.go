package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for i := 0; i < len(a); i++ {
		runeS := []rune(a[i])
		for j := 0; j < len(runeS); j++ {
			z01.PrintRune(runeS[j])
		}
		z01.PrintRune('\n')
	}
}
