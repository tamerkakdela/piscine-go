package piscine

func Join(strs []string, sep string) string {
	res := ""
	for index, i := range strs {
		res += i
		if index < len(strs)-1 {
			res += sep
		}
	}
	return res
}
