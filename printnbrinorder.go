package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
	}
	for n != 0 {
		s := (n % 10) + '0'
		z01.PrintRune(rune(s))
		n = n / 10
	}
}
