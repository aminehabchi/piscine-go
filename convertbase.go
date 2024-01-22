package piscine

import (
	"math"
)

func TenToBase(n int, baseTo int) []int {
	nb := []int{}
	for n != 0 {
		nb = append(nb, n%baseTo)
		n = n / baseTo
	}
	return nb
}

func ToTen(n []int, b int) int {
	sum := 0
	for i := len(n) - 1; i >= 0; i-- {
		sum = sum + n[i]*int(math.Pow(float64(b), float64(len(n)-1-i)))
	}
	return sum
}

func ConvertBase(nbr, baseFrom, baseTo string) string {
	if baseFrom == baseTo {
		return nbr
	}
	num := []int{}
	for i := 0; i < len(nbr); i++ {
		for j := 0; j < len(baseFrom); j++ {
			if nbr[i] == baseFrom[j] {
				num = append(num, j)
			}
		}
	}
	a := ToTen(num, len(baseFrom))
	b := TenToBase(a, len(baseTo))

	for i := 0; i < len(b)/2; i++ {
		b[i], b[len(b)-1-i] = b[len(b)-1-i], b[i]
	}
	str := ""
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(baseTo); j++ {
			if b[i] == j {
				str = str + string(baseTo[j])
			}
		}
	}
	return str
}
