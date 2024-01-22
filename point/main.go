package main

import (
	"github.com/01-edu/z01"
)

func rint(s int) {
	nb := '0'
	for i := 0; i < s/10; i++ {
		nb++
	}
	z01.PrintRune(nb)
	nb = '0'
	for i := 0; i < s%10; i++ {
		nb++
	}
	z01.PrintRune(nb)
}

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}

	setPoint(points)
	s1 := "x = a, y = b"
	for _, i := range s1 {
		if i == 'a' {
			rint(points.x)
		} else if i == 'b' {
			rint(points.y)
		} else {
			z01.PrintRune(i)
		}
	}

	z01.PrintRune('\n')
}
