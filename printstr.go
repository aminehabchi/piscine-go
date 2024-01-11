package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	a := []byte(s)

	for i := 0; i < len(s); i++ {
		z01.PrintRune(rune(a[i]))
	}
}
