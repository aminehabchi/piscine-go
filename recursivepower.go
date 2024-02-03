package piscine

func RecursivePower(nb int, power int) int {
	if power < 0 {
		return 0
	}

	result := 1
	if power > 0 {
		power--
		result = nb * RecursivePower(nb, power)
	}
	return result
}
