package piscine

func ConcatParams(args []string) string {
	var res string
	for i := 0; i <= len(args)-1; i++ {
		if i == len(args)-1 {
			res += args[i]
		} else {
			res += args[i] + "\n"
		}
	}
	return res
}
