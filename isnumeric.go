package piscine

func IsNumeric(s string) bool {
	ss := []rune(s)
	for i := 0; i < len(s); i++ {
		if ss[i] > '9' || ss[i] < '0' {
			return false
		}
	}
	return true
}
