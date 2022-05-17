package piscine

func IsAlpha(s string) bool {
	runeS := []rune(s)
	for _, str := range runeS {
		if str >= '0' && str <= '9' || str >= 'a' && str <= 'z' || str >= 'A' && str <= 'Z' {
			continue
		}
		return false
	}
	return true
}
