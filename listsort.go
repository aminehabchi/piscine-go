package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
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
