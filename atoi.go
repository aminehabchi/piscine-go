package piscine

func Atoi(s string) int {
	sum := 0
	count := 0
	sign := 1
	for _, num := range s {
		if num == '-' && count <= 1 {
			sign = sign * -1
			count++
			continue
		} else if num == '+' && count <= 1 {
			sign = sign * 1
			count++
			continue
		}

		if num >= 48 && num <= 57 && count <= 1 {
			sum = sum*10 + int(num-'0')
		} else {
			return 0
		}
	}
	sum = sign * sum
	return sum
}
