package piscine

type food struct {
	preptime int
}

func FoodDeliveryTime(order string) int {
	var burger food
	var chips food
	var nuggets food
	burger.preptime = 15
	chips.preptime = 10
	nuggets.preptime = 12
	switch {
	case order == "burger":
		return burger.preptime
	case order == "chips":
		return chips.preptime
	case order == "nuggets":
		return nuggets.preptime
	}
	return 0
}
