package piscine

func FindNextPrime(nb int) int {
	if nb <= 2 {
		return 2
	}

	var b bool
	var c float64
	for b == false {
		for i := 2; i < nb; i++ {
			c = float64(nb) / float64(i)
			if c == float64(int(c)) {
				b = false
				break
			}
			if c != float64(int(c)) {
				b = true
			}
		}
		if b == false {
			nb++
		}
	}
	return nb
}
