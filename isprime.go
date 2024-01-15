package piscine

func IsPrime(nb int) bool {
	var b bool
	var c float64
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
