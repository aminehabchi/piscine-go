package piscine
func ListPushFront(l *List, data interface{}) {
	n := &NodeL{Data: data}

	if l.Tail == nil {
		l.Tail = n
	} else {
		n.Next = l.Head
	}
	l.Head = n
}
type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}


func ListSize(l *List) int {
	c := l.Head
	a := 1
	for c.Next != nil {
		c = c.Next
		a++
	}
	return a
}
