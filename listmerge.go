package piscine

func ListMerge(l1 *List, l2 *List) {
	c := l1.Head
	if c == nil {
		l1.Head = l2.Head
	}
	for c != nil {
		if c.Next == nil {
			c.Next = l2.Head
			break
		}
		c = c.Next
	}
}
