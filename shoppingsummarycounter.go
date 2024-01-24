package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	s := Split(str, " ")
	m := make(map[string]int, len(s))
	for _, ele := range s {
		if _, ok := m[ele]; ok {
			m[ele]++
		} else {
			m[ele] = 1
		}
	}
	return m
}
