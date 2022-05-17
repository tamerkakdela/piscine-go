package piscine

func Index(s string, toFind string) int {
	if len(toFind) == 0 {
		return 0
	}
	if len(toFind) < len(s) {
		for i := 0; i < len(s); i++ {
			if s[i] == toFind[0] {
				counter := 0
				for j, k := 0, i; j < len(toFind); j, k = j+1, k+1 {
					if s[k] == toFind[j] {
						counter++
					}
					if counter == len(toFind) {
						return i
					}
				}
			}
		}
	}
	return -1
}
