package piscine

func ListSize(l *List) int {
	c := l.Head
	a := 1
	for c.Next != nil {
		c = c.Next
		a++
	}
	return a
}
