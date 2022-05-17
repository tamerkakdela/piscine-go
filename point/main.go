package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func IntToStr(a int) string {
	var astr []rune
	for a != 0 {
		tmp := a % 10
		astr = append(astr, rune(tmp+48))
		a = a / 10
	}
	var makeInverse string

	for i := len(astr) - 1; i >= 0; i-- {
		makeInverse += string(astr[i])
	}
	return makeInverse
}

func main() {
	points := &point{}

	setPoint(points)

	var printStr string

	printStr = "x = " + IntToStr(points.x) + ", y = " + IntToStr(points.y)

	for _, i := range printStr {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}
