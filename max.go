package piscine

func Max(a []int) int {
	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-i-1; j++ {
			if a[j+1] < a[j] {
				a[j+1], a[j] = a[j], a[j+1]
			}
		}
	}
	return a[len(a)-1]
}
