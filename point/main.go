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
	num := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	arr := "x = "

	for points.x > 0 {
		for i := 0; i < 10; i++ {
			if points.x%10 == i {
				arr = arr + num[i]
			}
		}
		points.x /= 10
	}
	arr = arr + ", y = "
	for points.y > 0 {
		for i := 0; i < 10; i++ {
			if points.y%10 == i {
				arr = arr + num[i]
			}
		}
		points.y /= 10
	}
	for _, char := range arr {
		z01.PrintRune(char)
	}

	z01.PrintRune('\n')
}
