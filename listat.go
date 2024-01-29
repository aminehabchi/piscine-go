package main

func ListAt(l *NodeL, pos int) *NodeL {
	a := 0
	c := l
	v := &NodeL
	for c != nil {
		if a == pos {
			return c
		}
		a++
		c = c.Next
	}
	return v
}
