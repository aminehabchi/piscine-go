package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]

	if len(arg) == 0 {
		for {
			c := ""
			fmt.Scanln(&c)
			for i := 0; i < len(c); i++ {
				z01.PrintRune(rune(c[i]))
			}
			z01.PrintRune('\n')
		}
	}

	for i := 0; i < len(arg); i++ {

		name := "./" + arg[i]
		content, err := ioutil.ReadFile(name)
		if err != nil && len(content) == 0 {
			print("ERROR: open " + name + ": no such file or directory")
		}
		for i := 0; i < len(content); i++ {
			z01.PrintRune(rune(content[i]))
		}

	}
}
