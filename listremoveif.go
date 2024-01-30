package piscine


func ListRemoveIf(l *List, data_ref interface{}) {
	c := l.Head
	var p *NodeL
	for c != nil {

		if c.Data == data_ref {
			if p == nil {
				l.Head = c.Next
			} else {
				p.Next = c.Next
			}
		} else {
			p = c
		}
		c = c.Next

	}
}
