package piscine

func TrimAtoi(s string) int {
	num := []rune{}
	ss := []rune(s)
	sign := 1
	b := false
	for i := 0; i < len(s); i++ {
		if ss[i] == '-' && b == false {
			sign = sign * -1
		}
		if s[i] <= '9' && ss[i] >= '0' {
			b = true
			num = append(num, ss[i])
		}
	}

	sum := 0
	for i := 0; i < len(num); i++ {
		sum = sum*10 + int(num[i])
	}

	return sum * sign
}
