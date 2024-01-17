package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrBase(nbr int, base string) {
	num := []int{}
	bn := len(base)
	ss := []rune(base)
	if nbr == 0 {
		z01.PrintRune(ss[0])
	}

	if len(base) < 2 {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}

	for i := 0; i < len(ss); i++ {
		for j := 0; j < len(ss); j++ {
			if j != i {
				if ss[i] == ss[j] || ss[i] == '-' || ss[i] == '+' {

					z01.PrintRune('N')
					z01.PrintRune('V')
					return
				}
			}
		}
	}

	if nbr < 0 {
		nbr = nbr * -1
		z01.PrintRune('-')
	}
	for nbr > 0 {
		num = append(num, nbr%bn)
		nbr /= bn
	}

	for i := 0; i < len(num); i++ {
		for j := i + 1; j < len(num); j++ {
			num[i], num[j] = num[j], num[i]
		}
	}

	for i := 0; i < len(num); i++ {
		z01.PrintRune(ss[num[i]])
	}
}
