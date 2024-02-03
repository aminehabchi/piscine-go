package piscine

func ToUpper(s string) string {
	ss := []rune(s)

	for i := 0; i < len(s); i++ {
		if ss[i] >= 97 && ss[i] <= 122 {
			ss[i] = rune(ss[i] - 32)
		}
	}
	return string(ss)
}
