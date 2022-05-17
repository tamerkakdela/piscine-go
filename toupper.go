package piscine

func ToUpper(s string) string {
	runeS := []rune(s)
	for i := 0; i < len(runeS); i++ {
		if runeS[i] >= 'a' && runeS[i] <= 'z' {
			runeS[i] -= 32
		}
	}
	return string(runeS)
}
