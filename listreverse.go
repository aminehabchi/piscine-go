package piscine

func ListReverse(l *List) {
	c := l.Head
	l.Tail = l.Head
	var p *NodeL
	var n *NodeL
	for c != nil {
		n = c.Next
		c.Next = p
		p = c
		c = n
	}
	l.Head = p
}
