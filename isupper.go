package piscine

func IsUpper(s string) bool {
	ss := []rune(s)
	for i := 0; i < len(s); i++ {
		if ss[i] > 'Z' || ss[i] < 'A' {
			return false
		}
	}
	return true
}
