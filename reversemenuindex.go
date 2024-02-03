package piscine

func ReverseMenuIndex(menu []string) []string {
	a := menu
	for i := len(menu) - 1; i >= 0; i-- {
		a[len(menu)-1-i] = menu[i]
	}
	return menu
}
