package piscine
func Atoi(s string) int {
	sum := 0
	sign := 1
	for index, num := range s {

		if num == '-' && index == 0 {
			sign = sign * -1
			continue
		} else if num == '+' && index == 0 {
			sign = sign * 1
			continue
		}

		if num >= 48 && num <= 57 {
			sum = sum*10 + int(num-'0')
		} else {
			return 0
		}

	}
	sum = sign * sum
	return sum
}
