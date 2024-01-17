package piscine

func Capitalize(s string) string {
	ss := []rune(s)
	for i := 0; i < len(s); i++ {
		if i == 0 {
			if ss[0] <= 'z' && ss[0] >= 'a' {
				ss[0] = ss[0] - 32
			}
			continue
		}
		if (ss[i] <= 'z' && ss[i] >= 'a') || (ss[i] <= 'Z' && ss[i] >= 'A') {
			if ss[i-1] < '0' || (ss[i-1] > '9' && ss[i-1] < 'A') || (ss[i-1] > 'Z' && ss[i-1] < 'a') || ss[i-1] > 'z' {
				if ss[i] <= 'z' && ss[i] >= 'a' {
					ss[i] = ss[i] - 32
				}
			} else if ss[i] <= 'Z' && ss[i] >= 'A' {
				ss[i] = ss[i] + 32
			}
		}
	}
	s = string(ss)
	return s
}
