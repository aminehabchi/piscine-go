package piscine

func BasicAtoi2(s string) int {
	sum := 0

	for _, num := range s {
		if num >= 48 && num <= 57 {
			sum = sum*10 + int(num-'0')
		} else {
			sum = 0
			break
		}
	}

	return sum
}
