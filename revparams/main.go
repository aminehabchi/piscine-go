package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argument := os.Args
	for i := len(argument) - 1; i > 0; i-- {
		b := os.Args[i]
		c := []rune(b)
		for j := 0; j < len(c); j++ {
			z01.PrintRune(c[j])
		}
		z01.PrintRune('\n')
	}
}
