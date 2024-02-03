package piscine

func AlphaCount(s string) int {
	count := 0
	for _, l := range s {
		if l >= 'a' && l <= 'z' || l >= 'A' && l <= 'Z' {
			count++
		}
	}

	return count
}
