package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := '0'; i <= '9'; i++ {
		for j := '0'; j <= '9'; j++ {
			d := j + 1
			for a := i; a <= '9'; a++ {
				for ; d <= '9'; d++ {
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(' ')
					z01.PrintRune(a)
					z01.PrintRune(d)
					if i < '9' || j < '8' {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}

				}
				d = '0'
			}
		}
	}
	z01.PrintRune('\n')

}
