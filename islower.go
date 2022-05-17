package piscine

func IsLower(s string) bool {
	runeS := []rune(s)
	for i := 0; i < len(runeS); i++ {
		if runeS[i] < 'a' || runeS[i] > 'z' {
			return false
		}
	}
	return true
}
