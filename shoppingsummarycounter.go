package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	a := map[string]int{"Burger": 0, "Water": 0, "Carrot": 0, "Coffee": 0, "Chips": 0}
	b := ""
	for i := 0; i < len(str); i++ {
		if str[i] != ' ' {
			b = b + string(str[i])
		}
		if str[i] == ' ' || i == len(str)-1 {
			switch {
			case b == "Burger":
				a["Burger"]++
			case b == "Water":
				a["Water"]++
			case b == "Carrot":
				a["Carrot"]++
			case b == "Coffee":
				a["Coffee"]++
			case b == "Chips":
				a["Chips"]++
			}
			b = ""
		}

	}

	return a
}
