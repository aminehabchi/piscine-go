package main

import (
	"bufio"
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]

	if len(arg) == 0 {
		for {
			reader := bufio.NewReader(os.Stdin)
			message, _ := reader.ReadString('\n')
			for i := 0; i < len(message); i++ {
				z01.PrintRune(rune(message[i]))
			}

		}
	}

	for i := 0; i < len(arg); i++ {
		name := "./" + arg[i]
		content, err := ioutil.ReadFile(name)
		if err != nil && len(content) == 0 {
			er := "ERROR: open " + arg[i] + ": no such file or directory"
			for i := 0; i < len(er); i++ {
				z01.PrintRune(rune(er[i]))
			}
			z01.PrintRune('\n')
			if len(arg) == 1 {
				os.Exit(1)
			}
		}
		for i := 0; i < len(content); i++ {
			z01.PrintRune(rune(content[i]))
		}

	}
}
