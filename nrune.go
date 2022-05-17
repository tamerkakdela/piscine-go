package piscine

func NRune(s string, n int) rune {
	if n <= 0 {
		return 0
	}
	res := []rune(s)

	if n > len(res) {
		return 0
	} else {
		return res[n-1]
	}
}
