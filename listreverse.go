package piscine

func ListReverse(l *List) {
	var next *NodeL = nil
	var prev *NodeL = nil
	cur := l.Head
	for cur != nil {
		next = cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	l.Head = prev
}
