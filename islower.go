package piscine

func IsLower(s string) bool {
	ss := []rune(s)
	for i := 0; i < len(s); i++ {
		if ss[i] > 'z' || ss[i] < 'a' {
			return false
		}
	}
	return true
}
