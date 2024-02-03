package piscine

func IsPrintable(s string) bool {
	ss := []rune(s)
	for i := 0; i < len(s); i++ {
		if ss[i] < ' ' {
			return false
		}
	}
	return true
}
