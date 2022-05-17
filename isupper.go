package piscine

func IsUpper(s string) bool {
	runeS := []rune(s)
	for i := 0; i < len(runeS); i++ {
		if runeS[i] < 'A' || runeS[i] > 'Z' {
			return false
		}
	}
	return true
}
