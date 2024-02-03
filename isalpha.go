package piscine

func IsAlpha(s string) bool {
	ss := []rune(s)
	for i := 0; i < len(s); i++ {
		if ss[i] < '0' || (ss[i] > '9' && ss[i] < 'A') || (ss[i] > 'Z' && ss[i] < 'a') || ss[i] > 'z' {
			return false
		}
	}
	return true
}
