package piscine

func ListMerge(l1 *List, l2 *List) {
	c := l1.Head
	if l1.Head == nil {
		l1.Head = l2.Head
		l1.Tail = l2.Tail
		return
	}
	for c != nil {
		if c.Next == nil {
			c.Next = l2.Head
			break
		}
		c = c.Next
	}
}
