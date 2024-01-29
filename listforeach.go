package piscine

func ListForEach(l *List, f func(*NodeL)) {
	c := l.Head

	for c != nil {
		f(c)
		c = c.Next
	}
}
