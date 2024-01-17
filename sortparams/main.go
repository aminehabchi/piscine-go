package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argument := os.Args
	ar := []rune{}
	for i := 1; i < len(argument); i++ {

		c := os.Args[i]
		d := []rune(c)
		ar = append(ar, d[0])

	}
	for i := 0; i < len(ar); i++ {
		for j := i + 1; j < len(ar); j++ {
			if j != i {
				if ar[i] > ar[j] {
					ar[i], ar[j] = ar[j], ar[i]
				}
			}
		}
	}
	for i := 0; i < len(ar); i++ {
		z01.PrintRune(ar[i])
		z01.PrintRune('\n')
	}
}
