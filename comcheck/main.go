package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	for i := 0; i < len(arg); i++ {
		if arg[i] == "galaxy" || arg[i] == "galaxy 01" || arg[i] == "01" {
			print("Alert!!!")
			z01.PrintRune('A')
			z01.PrintRune('l')
			z01.PrintRune('e')
			z01.PrintRune('r')
			z01.PrintRune('t')
			z01.PrintRune('!')
			z01.PrintRune('!')
			z01.PrintRune('!')
			break
		}
	}
}
