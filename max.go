package piscine

func Max(a []int) int {
	b := a[0]
	for i := 0; i < len(a); i++ {
		if b < a[i] {
			b = a[i]
		}
	}
	return b
}
