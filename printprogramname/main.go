package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	b := ""
	for i := 0; i < len(arguments); i++ {
		b = b + arguments[i]
	}
	c := []rune(b)
	for i := 2; i < len(c); i++ {
		z01.PrintRune(c[i])
	}
	z01.PrintRune('\n')
}
