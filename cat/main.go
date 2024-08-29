package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	a := os.Args[1:]

	arr := make([]byte, 1024)
	exit := 1
	if len(a) == 0 {
		da, err := os.Stdin.Read(arr)
		if err == nil {
			PrintStr(string(arr[:da]))
		}
	}
	for i := 0; i < len(a); i++ {
		da, err := ioutil.ReadFile(a[i])
		if err != nil {
			PrintStr("ERROR: open " + a[i] + ": no such file or directory\n")
			os.Exit(exit)
			return
		}
		PrintStr(string(da))
	}
}

func PrintStr(v string) {
	for i := 0; i < len(v); i++ {
		z01.PrintRune(rune(v[i]))
	}
}
