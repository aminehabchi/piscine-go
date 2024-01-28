package piscine

func AtoiBase(s string, base string) int {
	ss := []rune(s)
	bb := []rune(base)
	b := len(base)
	num := []int{}
	sum := 0
	// If the base is not valid it returns 0.
	if len(s) < 2 {
		return 0
	}
	for i := 0; i < len(bb); i++ {
		for j := 0; j < len(bb); j++ {
			if j != i {
				if bb[i] == bb[j] || bb[i] == '-' || bb[i] == '+' {
					return 0
				}
			}
		}
	}

	// String number must contain only elements that are in base.

	for i := 0; i < len(ss); i++ {
		bol := false
		for j := 0; j < len(bb); j++ {
			if ss[i] == bb[j] {
				bol = true
			}
		}
		if bol == false {
			return 0
		}
	}

	for i := 0; i < len(ss); i++ {
		for j := 0; j < len(bb); j++ {
			if ss[i] == bb[j] {
				num = append(num, j)
			}
		}
	}

	for i := 0; i < len(num); i++ {
		sum = sum*b + num[i]
	}

	return sum
}
