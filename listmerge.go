package piscine

func ListMerge(l1 *List, l2 *List) {
	c := l1.Head
	for c != nil {
		if c.Next == nil {
			c.Next = l2.Head
			break
		}
		c = c.Next
	}
}
