package piscine

func IsNumeric(s string) bool {
	runeS := []rune(s)
	for _, str := range runeS {
		if str >= '0' && str <= '9' {
			continue
		}
		return false
	}
	return true
}
