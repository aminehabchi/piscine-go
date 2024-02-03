package piscine

func IsPrime(nb int) bool {
	var b bool
	var c float64
	if nb == 2 {
		return true
	}
	if nb == 1 {
		return false
	}
	for i := 2; i < nb; i++ {
		c = float64(nb) / float64(i)
		if c == float64(int(c)) {
			return false
		}
		if c != float64(int(c)) {
			b = true
		}

	}
	return b
}
