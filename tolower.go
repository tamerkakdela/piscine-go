package piscine

func ToLower(s string) string {
	runeS := []rune(s)
	for i := 0; i < len(runeS); i++ {
		if runeS[i] >= 'A' && runeS[i] <= 'Z' {
			runeS[i] = runeS[i] + 32
		}
	}
	return string(runeS)
}
