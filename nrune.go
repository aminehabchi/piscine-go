package piscine

func NRune(s string, n int) rune {
	if n < 1 || n > len(s) {
		return 0
	}
	b := []rune(s)
	return b[n-1]
}
