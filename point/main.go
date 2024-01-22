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
	num := []rune{48, 49, 50, 51, 52, 53, 54, 56, 57}
	//"x = %, y = %\n"
	for i := 0; i < 12; i++ {
		c := 0
		if i == 0 {
			z01.PrintRune('x')
		}
		if i == 7 {
			z01.PrintRune('y')
		}
		if i == 4 || i == 11 {
			if i == 4 {
				c = points.x
			} else {
				c = points.y
			}
			for c > 0 {
				for i := 0; i < 10; i++ {
					if c%10 == i {
						z01.PrintRune(num[i])
					}
				}
				c /= 10
			}
		}
		if i == 1 || i == 3 || i == 6 || i == 8 || i == 10 {
			z01.PrintRune(' ')
		}
		if i == 2 || i == 9 {
			z01.PrintRune('=')
		}
		if i == 5 {
			z01.PrintRune(',')
		}
	}

	z01.PrintRune('\n')
}
