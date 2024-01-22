package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	if even(nbr) == true {
		return true
	}
	return false
}

func even(nbr int) bool {
	if nbr%2 == 0 {
		return true
	}
	return false
}

func main() {
	arg := os.Args[1:]
	lengthOfArg := len(arg)
	EvenMsg := "I have an even number of arguments"
	OddMsg := "I have an odd number of arguments"
	if isEven(lengthOfArg) == true {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}
