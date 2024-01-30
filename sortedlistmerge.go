package piscine

func Listmerge(n1 *NodeI, n2 *NodeI) *NodeI {
	if n1 == nil {
		return n2
	}
	if n2 == nil {
		return n1
	}

	current := n1
	for current.Next != nil {
		current = current.Next
	}
	current.Next = n2

	return n1
}

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	v := Listmerge(n1, n2)
	a := Listsort(v)
	return a
}

func Listsort(l *NodeI) *NodeI {
	a := l
	c := 0
	for a != nil {
		a = a.Next
		c++

	}

	for i := 0; i < c; i++ {
		a = l
		for a.Next != nil {
			if a.Data > a.Next.Data {
				a.Data, a.Next.Data = a.Next.Data, a.Data
			}
			a = a.Next
		}
	}
	return l
}
