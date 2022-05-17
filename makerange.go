package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	r := 0
	result := make([]int, max-min)
	for i := min; i < max; i++ {
		result[r] = i
		r++
	}
	return result
}
