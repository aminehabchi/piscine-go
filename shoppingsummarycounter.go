package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	m := map[string]int{}
	s := ""
	for index, char := range str {
		if char == ' ' || index == len(str)-1 {
			m[s] = m[s] + 1
			s = ""
		} else {
			s = s + string(char)
		}
	}
	return m
}
