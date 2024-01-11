package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	a := '0'
	b := '0'
	c := '0'
	d := '0'
	e := '0'
	f := '0'
	g := '0'
	h := '0'
	i := '0'

	for i < '9' {

		if n == 1 {
			z01.PrintRune(a)
			if a != '9' {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
		if n == 2 && a > b {
			z01.PrintRune(b)
			z01.PrintRune(a)
			if b != '8' {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}

		}
		if n == 3 && a > b && b > c {
			z01.PrintRune(c)
			z01.PrintRune(b)
			z01.PrintRune(a)
			if c != '7' {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}

		}
		if n == 4 && a > b && b > c && c > d {
			z01.PrintRune(d)
			z01.PrintRune(c)
			z01.PrintRune(b)
			z01.PrintRune(a)
			if d != '6' {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}

		}
		if n == 5 && a > b && b > c && c > d && d > e {
			z01.PrintRune(e)
			z01.PrintRune(d)
			z01.PrintRune(c)
			z01.PrintRune(b)
			z01.PrintRune(a)
			if e != '5' {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}

		}
		if n == 6 && a > b && b > c && c > d && d > e && e > f {

			z01.PrintRune(f)
			z01.PrintRune(e)
			z01.PrintRune(d)
			z01.PrintRune(c)
			z01.PrintRune(b)
			z01.PrintRune(a)
			if f != '4' {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
		if n == 7 && a > b && b > c && c > d && d > e && e > f && f > g {
			z01.PrintRune(g)
			z01.PrintRune(f)
			z01.PrintRune(e)
			z01.PrintRune(d)
			z01.PrintRune(c)
			z01.PrintRune(b)
			z01.PrintRune(a)
			if g != '3' {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}

		}
		if n == 8 && a > b && b > c && c > d && d > e && e > f && f > g && g > h {

			z01.PrintRune(g)
			z01.PrintRune(f)
			z01.PrintRune(e)
			z01.PrintRune(d)
			z01.PrintRune(c)
			z01.PrintRune(b)
			z01.PrintRune(a)
			if h != '3' {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
		if n == 9 && a > b && b > c && c > d && d > e && e > f && f > g && g > h && h > i {

			z01.PrintRune(i)
			z01.PrintRune(h)
			z01.PrintRune(g)
			z01.PrintRune(f)
			z01.PrintRune(e)
			z01.PrintRune(d)
			z01.PrintRune(c)
			z01.PrintRune(b)
			z01.PrintRune(a)
			if i != '3' {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}

		if a < '9' {
			a++
		} else if b < '9' {

			if n == 1 && a == '9' {
				break
			}

			b++
			a = '0'
		} else if c < '9' {
			if n == 2 && a == '9' && b == '9' {
				break
			}
			c++
			b = '0'
			a = '0'
		} else if d < '9' {
			if n == 3 && a == '9' && b == '9' && c == '9' {
				break
			}
			d++
			c = '0'
			b = '0'
			a = '0'
		} else if e < '9' {
			if n == 4 && a == '9' && b == '9' && c == '9' && d == '9' {
				break
			}
			e++
			d = '0'
			c = '0'
			b = '0'
			a = '0'
		} else if f < '9' {
			if n == 5 && a == '9' && b == '9' && c == '9' && d == '9' && e == '9' {
				break
			}
			f++
			e = '0'
			d = '0'
			c = '0'
			b = '0'
			a = '0'
		} else if g < '9' {
			if n == 6 && a == '9' && b == '9' && c == '9' && d == '9' && e == '9' && f == '9' {
				break
			}
			g++
			f = '0'
			e = '0'
			d = '0'
			c = '0'
			b = '0'
			a = '0'
		} else if h < '9' {
			if n == 7 && a == '9' && b == '9' && c == '9' && d == '9' && e == '9' && f == '9' && g == '9' {
				break
			}
			h++
			g = '0'
			f = '0'
			e = '0'
			d = '0'
			c = '0'
			b = '0'
			a = '0'
		} else if i < '9' {
			if n == 8 && a == '9' && b == '9' && c == '9' && d == '9' && e == '9' && f == '9' && g == '9' && h == '9' {
				break
			}
			i++
			h = '0'
			g = '0'
			f = '0'
			e = '0'
			d = '0'
			c = '0'
			b = '0'
			a = '0'
		}
		if n == 9 && a == '9' && b == '9' && c == '9' && d == '9' && e == '9' && f == '9' && g == '9' && h == '9' && i == '9' {
			break
		}

	}
	z01.PrintRune('\n')
}
