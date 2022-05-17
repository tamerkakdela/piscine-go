package piscine

func Capitalize(s string) string {
	runeS := []rune(s)
	for i := 0; i < len(runeS); i++ {
		if runeS[i] >= 'A' && runeS[i] <= 'Z' {
			runeS[i] += 32
		}
	}
	count := 0
	for i := 0; i < len(runeS); i++ {
		if IsAlphaA(runeS[i]) == true {
			count++
			if count == 1 {
				if runeS[i] < '0' || runeS[i] > '9' {
					runeS[i] -= 32
				}
			}

		}
		if IsAlphaA(runeS[i]) == false {
			count = 0
		}
	}
	return string(runeS)
}

func IsAlphaA(runeS rune) bool {
	if runeS >= 'a' && runeS <= 'z' || runeS >= 'A' && runeS <= 'Z' || runeS >= '0' && runeS <= '9' {
		return true
	}
	return false
}
