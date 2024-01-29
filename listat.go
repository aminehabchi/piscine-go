package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	a := 0
	c := l

	for c != nil {
		if a == pos {
			return c
		}
		a++
		c = c.Next
	}
	c = nil
	return c
}
