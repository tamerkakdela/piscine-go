package piscine

func IsPrintable(s string) bool {
	runeS := []rune(s)
	count := 0
	for _, str := range runeS {
		if str == '\n' || str == '\r' || str == '\a' || str == '\v' || str == '\b' {
			count++
		}
		if count != 0 {
			return false
		}
	}
	return true
}
