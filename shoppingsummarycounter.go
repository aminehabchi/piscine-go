package main

import (
	"fmt"
)

func ShoppingSummaryCounter(str string) map[string]int {
	a := map[string]int{"Burger": 0, "Water": 0, "Carrot": 0, "Coffee": 0, "Chips": 0}
	b := ""
	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {
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
		} else {
			b = b + string(str[i])
		}
	}

	return a
}

func main() {
	summary := "Burger Water Carrot Coffee Water Water Chips Carrot Carrot Burger Carrot Water"
	for index, element := range ShoppingSummaryCounter(summary) {
		fmt.Println(index, "=>", element)
	}
}
