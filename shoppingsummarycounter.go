package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	a := map[string]int{}
	b := ""
	bb := []string{}
	for i := 0; i < len(str); i++ {
		if str[i] != ' ' {
			b = b + string(str[i])
		}
		if str[i] == ' ' || i == len(str)-1 {
			bb = append(bb, b)
			b = ""
		}
	}
	for i := 0; i < len(bb); i++ {
		for j := 0; j < len(bb); j++ {
			if bb[i] > bb[j] {
				bb[i], bb[j] = bb[j], bb[i]
			}
		}
	}

	for i := 1; i < len(bb); i++ {
		if bb[i] != bb[i-1] {
			a[bb[i]] = 0
		}
	}

	for i := 0; i < len(bb); i++ {
		count := 0
		for j := 0; j < len(bb); j++ {
			if bb[i] == bb[j] {
				count++
			}
		}
		if bb[i] == "" {
			count++
		}
		a[bb[i]] = count
	}

	return a
}
