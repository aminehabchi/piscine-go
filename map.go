package piscine

func Map(f func(int) bool, a []int) []bool {
	b := []bool{}
	for i := 0; i < len(a); i++ {
		b = append(b, f(a[i]))
	}
	return b
}
