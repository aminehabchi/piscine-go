package piscine

func ToLower(s string) string {
	ss := []rune(s)

	for i := 0; i < len(s); i++ {
		if ss[i] >= 65 && ss[i] <= 90 {
			ss[i] = rune(ss[i] + 32)
		}
	}
	return string(ss)
}
