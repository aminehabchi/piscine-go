package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	n := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = n
	}
	if l.Head != nil {
		c := l.Head
		if c.Next != nil {
			c = c.Next
		}
		c.Next = n
	}
}
