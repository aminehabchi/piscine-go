package piscine

func JumpOver(str string) string {
	n := "\n"
	if str == "" || len(str) < 3 {
		return n
	}
	s := ""
	a := 0
	for i := 0; i < len(str); i++ {
		if a == 2 {
			s = s + string(str[i])
			a = 0
		} else {
			a++
		}
	}
	s = s + n
	return s
}
