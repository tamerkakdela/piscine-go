package piscine

func Rot14(s string) string {
	runeS := []rune(s)
	for i := 0; i < len(runeS); i++ {
		if runeS[i] >= 'A' && runeS[i] <= 'L' || runeS[i] >= 'a' && runeS[i] <= 'l' {
			runeS[i] += 14
		} else if runeS[i] >= 'M' && runeS[i] <= 'Z' || runeS[i] >= 'm' && runeS[i] <= 'z' {
			runeS[i] -= 12
		}
	}
	return string(runeS)
}
