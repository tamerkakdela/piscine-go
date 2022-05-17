package piscine

func SplitWhiteSpaces(s string) []string {
	res := []rune(s)
	var sep []string
	word := ""
	for i := 0; i < len(res); i++ {
		if !(res[i] == 32 || res[i] == 9 || res[i] == 8) {
			word += string(res[i])
		} else {
			if word != "" {
				sep = append(sep, word)
			}
			word = ""
		}
	}
	if word != "" {
		sep = append(sep, word)
	}
	return sep
}
