package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func PrintNbr(n int) {
	t := 1
	if n < 0 {
		t = -1
		z01.PrintRune('-')
	}
	if n != 0 {
		f := (n / 10) * t
		if f != 0 {
			PrintNbr(f)
		}
		k := (n % 10 * t) + '0'
		z01.PrintRune(rune(k))
	} else {
		z01.PrintRune('0')
	}
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}

	setPoint(points)

	p := "x = ., y = :\n"
	for i := 0; i < len(p); i++ {
		if p[i] == '.' {
			PrintNbr(points.x)
		} else if p[i] == ':' {
			PrintNbr(points.y)
		} else {
			z01.PrintRune(rune(p[i]))
		}
	}
}
