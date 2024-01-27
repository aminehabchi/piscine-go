package main

import (
	"os"

	"github.com/01-edu/z01"
)

func Atoi(s string) (uint, bool) {
	var sum uint = 0
	signe := 1
	b := false
	for index, num := range s {

		if num == '-' && index == 0 {
			signe = signe * -1
			continue
		} else if num == '+' && index == 0 {
			signe = signe * 1
			continue
		}

		if num >= 48 && num <= 57 {
			sum = sum*10 + uint(num-'0')
		} else {
			return 0, b
		}

	}

	if signe == 1 && sum > 9223372036854775807 {
		return 0, b
	}
	if signe == -1 && sum > 9223372036854775808 {
		return 0, b
	}

	if (signe == 1 && sum == 9223372036854775807) || (signe == -1 && sum == 9223372036854775808) {
		b = true
	}

	sum = uint(signe) * sum
	return sum, b
}

func PrintNbr(n int) {
	t := 1
	if n < 0 {
		t = -1
		z01.PrintRune('-')
	}
	if n != 0 {
		f := (n / 10) * t
		if f != 0 {
			PrintNbr(f)
		}
		k := (n % 10 * t) + '0'
		z01.PrintRune(rune(k))
	} else {
		z01.PrintRune('0')
	}
	z01.PrintRune('\n')
}

func main() {
	arg := os.Args[1:]
	if len(arg) != 3 {
		return
	}
	if arg[1] != "+" && arg[1] != "-" && arg[1] != "*" && arg[1] != "/" && arg[1] != "%" {
		return
	}
	num1, b1 := Atoi(arg[0])
	num2, b2 := Atoi(arg[2])
	if num1 == 0 && arg[0] != "0" {
		return
	}
	if num1 == 0 && arg[2] != "0" {
		return
	}
	if arg[1] == "+" {
		if b1 == true || b2 == true {
			return
		}

		a := int(num1 + num2)
		PrintNbr(a)
	}
	if arg[1] == "-" {
		a := int(num1 - num2)
		PrintNbr(a)
	}
	if arg[1] == "*" {
		if b1 == true || b2 == true {
			return
		}
		a := int(num1 * num2)
		PrintNbr(a)
	}
	if arg[1] == "/" {
		if num2 == 0 {
			print("No division by 0")
			return
		}
		a := int(num1 / num2)
		PrintNbr(a)
	}
	if arg[1] == "%" {
		if num2 == 0 {
			print("No modulo by 0")
			return
		}
		a := int(num1 % num2)
		PrintNbr(a)
	}
}
