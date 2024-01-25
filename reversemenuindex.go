package piscine

func ReverseMenuIndex(menu []string) []string {
	a := menu
	for i := len(menu) - 1; i >= 0; i-- {
		a = menu[i]
	}
	return menu
}
