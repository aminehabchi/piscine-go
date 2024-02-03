package piscine

func SplitWhiteSpaces(s string) []string {
	ss := []rune(s)
	b := ""
	str := []string{}
	for i := 0; i < len(ss); i++ {
		if ss[i] != ' ' {
			b = b + string(ss[i])
		}
		if (ss[i] == ' ' || i == len(ss)-1) && b != "" {
			str = append(str, b)
			b = ""
		}
	}
	return str
}
