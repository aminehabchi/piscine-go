package piscine

func Index(s string, toFind string) int {
	if len(s) < len(toFind) {
		return -1
	}

	if s == toFind {
		return 0
	}
	if len(toFind) == 1 {
		for index, char := range s {
			if string(char) == toFind {
				return index
			}
		}
	}
	sr := []rune(s)
	fr := []rune(toFind)
	var b string

	for i := 0; i < len(s)-len(toFind); i++ {
		if sr[i] == fr[0] {
			b = ""
			for j := 0; j < len(toFind); j++ {
				b = b + string(sr[i+j])
			}
			if b == toFind {
				return i
			}
		}
	}

	return -1
}
