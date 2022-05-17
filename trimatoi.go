package piscine

func TrimAtoi(s string) int {
	runeS := []rune(s)

	var sInt []rune
	var number int
	var minusSign bool = false

	for i := 0; i < len(runeS); i++ {
		if runeS[i] >= 48 && runeS[i] <= 57 {
			for j := 0; j < i; j++ {
				if runeS[j] == '-' {
					minusSign = true
				}
			}
			break
		}
	}

	for i := 0; i < len(runeS); i++ {
		if runeS[i] >= 48 && runeS[i] <= 57 {
			sInt = append(sInt, runeS[i])
		}
	}
	for i, j := len(sInt)-1, 1; i >= 0; i, j = i-1, j*10 {
		number = int(sInt[i]-'0')*j + number
	}
	if minusSign {
		return -1 * number
	}
	return number
}
