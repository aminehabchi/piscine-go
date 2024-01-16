package piscine

func Compare(a, b string) int {
	aa := []rune(a)
	bb := []rune(b)
	if aa[0] != bb[0] {
		return -1
	}

	if len(a) == len(b) {
		x := 0
		for i := 0; i < len(b); i++ {
			if aa[i] == bb[i] {
				x = 0
			} else {
				return -1
			}
		}
		return x
	}
	return 1
}
