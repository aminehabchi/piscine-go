package piscine

func RockAndRoll(n int) string {
	b := "number is negative\n"
	if n < 0 {
		return b
	}
	b = "non divisible\n"
	if n%2 == 0 && n%3 == 0 {
		b = "rock and roll\n"
	} else if n%2 == 0 {
		b = "rock\n"
	} else if n%3 == 0 {
		b = "roll\n"
	}

	return b
}
