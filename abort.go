package piscine

func Abort(a, b, c, d, e int) int {
	var slice []int
	slice = append(slice, a, b, c, d, e)
	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-i-1; j++ {
			if slice[j+1] < slice[j] {
				slice[j+1], slice[j] = slice[j], slice[j+1]
			}
		}
	}
	return slice[2]
}
