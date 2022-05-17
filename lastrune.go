package piscine

func LastRune(s string) rune {
	res := []rune(s)
	newRes := res[len(res)-1]
	return newRes
}
