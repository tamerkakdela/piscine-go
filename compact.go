package piscine

func Compact(ptr *[]string) int {
	count := 0
	var slice []string
	for _, i := range *ptr {
		if i != "" {
			slice = append(slice, i)
			count++
		}
	}
	*ptr = slice
	return count
}
