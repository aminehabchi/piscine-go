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

func main() {
	points := &point{}

	setPoint(points)
	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	num := []rune{48, 49, 50, 51, 52, 53, 54, 56, 57}

	for points.x > 0 {
		for i := 0; i < 10; i++ {
			if points.x%10 == i {
				z01.PrintRune(num[i])
			}
		}
		points.x /= 10
	}

	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune('=')
	z01.PrintRune(' ')

	for points.y > 0 {
		for i := 0; i < 10; i++ {
			if points.y%10 == i {
				z01.PrintRune(num[i])
			}
		}
		points.y /= 10
	}
	z01.PrintRune('\n')
}
