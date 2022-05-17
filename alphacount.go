package piscine

func AlphaCount(s string) int {
	count := 0
	for _, str := range s {
		if str >= 'a' && str <= 'z' || str >= 'A' && str <= 'Z' {
			count++
		}
	}
	return count
}
