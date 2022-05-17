package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	Up := false
	Down := false
	for i := 0; i < len(a)-1; i++ {
		if Up {
			if f(a[i], a[i+1]) > 0 {
				return false
			}
		}
		if Down {
			if f(a[i], a[i+1]) < 0 {
				return false
			}
		}

		if !Up && !Down {
			if f(a[i], a[i+1]) > 0 {
				Down = true
			}
			if f(a[i], a[i+1]) < 0 {
				Up = true
			}
		}
	}
	return true
}
