package piscine

func Any(f func(string) bool, a []string) bool {
	b := []bool{}
	for i := 0; i < len(a); i++ {
		b = append(b, f(a[i]))
	}
	for i := 0; i < len(b); i++ {
		if b[i] == true {
			return true
		}
	}
	return false
}
