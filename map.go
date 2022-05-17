package piscine

func Map(f func(int) bool, a []int) []bool {
	var answer []bool
	for i := 0; i < len(a); i++ {
		answer = append(answer, f(a[i]))
	}
	return answer
}
