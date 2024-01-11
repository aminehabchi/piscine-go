package piscine

func BasicAtoi(s string) int {
	sum := 0

	for _, num := range s {
		sum = sum*10 + int(num-'0')
	}

	return sum
}
