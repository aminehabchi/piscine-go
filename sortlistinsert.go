package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	b := &NodeI{Data: data_ref}
	b.Next = l
	b = piscine.ListSort(b)
	return b
}
