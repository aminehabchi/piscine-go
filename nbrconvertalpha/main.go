package main

import (
	"os"

	"github.com/01-edu/z01"
)

func Atoi(s string) int {
	sum := 0
	signe := 1
	for index, num := range s {

		if num == '-' && index == 0 {
			signe = signe * -1
			continue
		} else if num == '+' && index == 0 {
			signe = signe * 1
			continue
		}

		if num >= 48 && num <= 57 {
			sum = sum*10 + int(num-'0')
		} else {
			return 0
		}

	}
	sum = signe * sum
	return sum
}

func main() {
	argiment := os.Args
	a := 96
	if os.Args[1] == "--upper" {
		a = 64
	}
	for i := 2; i < len(argiment); i++ {

		b := Atoi(os.Args[i])
		if b < 1 || b > 26 {
			z01.PrintRune(' ')
			continue
		}
		z01.PrintRune(rune(b + a))
	}
	z01.PrintRune('\n')
}
