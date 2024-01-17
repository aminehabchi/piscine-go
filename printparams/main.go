package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argiment := os.Args

	for i := 1; i < len(argiment); i++ {
		b := os.Args[i]
		c := []rune(b)
		for j := 0; j < len(c); j++ {
			z01.PrintRune(c[j])
		}
		z01.PrintRune('\n')
	}
}
