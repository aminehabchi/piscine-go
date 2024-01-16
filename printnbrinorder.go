package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
	}

	ss := []int{}

	for n > 0 {
		ss = append(ss, n%10)
		n = n / 10
	}

	for i := 0; i < len(ss); i++ {
		for j := i + 1; j < len(ss); j++ {
			if ss[i] > ss[j] {
				ss[i], ss[j] = ss[j], ss[i]
			}
		}
	}

	for i := 0; i < len(ss); i++ {
		z01.PrintRune(rune(ss[i] + '0'))
	}
}
