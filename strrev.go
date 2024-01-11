package piscine

func StrRev(s string) string {
	a := []byte(s)
	b := []byte(s)

	for i := 0; i < len(a); i++ {
		b[len(a)-i-1] = a[i]
	}
	s = string(b)
	return s
}
