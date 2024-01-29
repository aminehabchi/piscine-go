package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	a := 0
	c := l
	for c != nil {
		a++
		c = c.Next
		if a == pos {
			return c
		}
	}
	return c
}
