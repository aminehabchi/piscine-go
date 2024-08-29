package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argument := os.Args

	for i := 1; i < len(argument); i++ {
		for j := 1; j < len(argument); j++ {
			if os.Args[i] < os.Args[j] {
				if i != j {
					os.Args[i], os.Args[j] = os.Args[j], os.Args[i]
				}
			}
		}
	}
	for i := 1; i < len(argument); i++ {
		b := []rune(os.Args[i])

		for j := 0; j < len(b); j++ {
			z01.PrintRune(b[j])
		}
		z01.PrintRune('\n')
	}
}
