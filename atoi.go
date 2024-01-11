package main

import (
	"fmt"
)

func Atoi(s string) int {
	sum := 0
	count := 0
	sign := 1
	for _, num := range s {
		if num == 45 && count <= 1 {
			sign = sign * -1
			count++
			continue
		} else if num == 43 && count <= 1 {
			sign = sign * 1
			count++
			continue
		}

		if num >= 48 && num <= 57 && count <= 1 {
			sum = sum*10 + int(num-'0')
		} else {
			return 0
		}
	}
	sum = sign * sum
	return sum
}

func main() {
	fmt.Println(Atoi("12345"))
	fmt.Println(Atoi("0000000012345"))
	fmt.Println(Atoi("012 345"))
	fmt.Println(Atoi("Hello World!"))
	fmt.Println(Atoi("+1234"))
	fmt.Println(Atoi("-1234"))
	fmt.Println(Atoi("++1234"))
	fmt.Println(Atoi("--1234"))
}
