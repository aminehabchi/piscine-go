package piscine

func RecursiveFactorial(nb int) int {
	n := nb
	var rusult int = n

	if nb > 1 {
		n = n - 1
		rusult = rusult * RecursiveFactorial(n)
	}
	if nb == 0 {
		return 1
	}
	if nb < 0 || nb > 31 {
		return 0
	}

	return rusult
}
