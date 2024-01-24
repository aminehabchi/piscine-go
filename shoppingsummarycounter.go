package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	s := Split(str, " ")
	a := make(map[string]int, len(s))
	for _, ele := range s {
		if _, ok := a[ele]; ok {
			a[ele]++
		} else {
			a[ele] = 1
		}
	}
	return a
}
