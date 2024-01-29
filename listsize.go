package piscine

func ListSize(l *List) int {
	c := l.Head
	a := 0
	for c != nil {
		c = c.Next
		a++
	}
	return a
}
