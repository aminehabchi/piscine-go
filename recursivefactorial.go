package piscine

func RecursiveFactorial(nb int) int {
	if nb == 0 {
		return 1
	}
	if nb < 0 || nb > 31 {
		return 0
	}
	n := nb
	var rusult int = n
	if nb > 1 {
		n = n - 1
		rusult = rusult * RecursiveFactorial(n)
	}
	return rusult
}
