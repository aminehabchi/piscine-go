package main

import (
	"fmt"
)

func LoafOfBread(str string) string {
	b := ""
	for i := 0; i < len(str); i++ {
		if str[i] != ' ' {
			b = b + string(str[i])
		}
	}
	if len(b) < 5 {
		b = "Invalid Output\n"
		return b
	}
	a := 5
	d := ""
	for i := 0; i < len(b); i++ {

		d = d + string(b[i])
		if i == a {
			d = d + " "
			a += 5
		}

	}
	println(b)
	d = d + "\n"
	return d
}

func main() {
	fmt.Print(LoafOfBread("deliciousbread"))
	fmt.Print(LoafOfBread("This is a loaf of bread"))
	fmt.Print(LoafOfBread("loaf"))
}
