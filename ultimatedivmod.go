package piscine

func UltimateDivMod(a *int, b *int) {
	d := *a / *b
	m := *a % *b
	*a = d
	*b = m
}
