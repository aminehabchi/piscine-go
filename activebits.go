package piscine

func ActiveBits(n int) int {
	nn := []int{}
	for n != 0 {
		nn = append(nn, n%2)
		n /= 2
	}
	return len(nn)
}
