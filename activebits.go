package piscine

func ActiveBits(n int) int {
	nn := []int{}
	for n != 0 {
		nn = append(nn, n%2)
		n /= 2
	}
	c := 0
	for i := 0; i < len(nn); i++ {
		if nn[i] == 1 {
			c++
		}
	}

	return c
}
