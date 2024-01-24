package piscine

func Compact(ptr *[]string) int {
	a := *ptr
	b := []string{}
	count := 0
	for i := 0; i < len(a); i++ {
		if a[i] != "" {
			b = append(b, a[i])
		}
	}
	*ptr = b
	return count
}
