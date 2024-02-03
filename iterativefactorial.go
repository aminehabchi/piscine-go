package piscine

func IterativeFactorial(nb int) int {
	if nb == 0 {
		return 1
	}
	if nb < 0 || nb > 31 {
		return 0
	}
	s := 1
	for i := 1; i <= nb; i++ {
		s = i * s
	}
	return s
}
