package piscine

func AppendRange(min, max int) []int {
	if min > max {
		return nil
	}

	var str []int
	for i := min; i < max; i++ {
		str = append(str, min)
		min++
	}
	return str
}
