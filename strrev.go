package piscine

func StrRev(s string) string {
	a := []rune(s)
	length := len(a)
	arr := make([]rune, length)
	for i := 0; i < length; i++ {
		arr[i] = a[length-i-1]
	}

	return string(arr)
}
