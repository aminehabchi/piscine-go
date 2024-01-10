package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbr(n int) {
	if n == 0 {
		z01.PrintRune('0')
	}
	if n < 0 {
		z01.PrintRune('-')
		n = n * -1
	}
	c := 1
	i := 0
	p := 0
	// length i c=1000
	for c > 0 {

		if n/c <= 0 {
			break
		}
		i++
		c = c * 10
	}
	c = c / 10
	var a int
	var b int
	for j := 0; j < i; j++ {
		a = n % c
		b = n - a
		n = a
		b = b / c
		c = c / 10
		for k := '0'; k <= '9'; k++ {
			if p == b {
				z01.PrintRune(k)
			}

			if p == 9 {
				p = 0
			} else {
				p++
			}
		}

	}
}
