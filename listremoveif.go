package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	curr := l.Head
	var prev *NodeL

	for curr != nil {

		if curr.Data == data_ref {
			if prev == nil {
				l.Head = curr.Next
			} else {
				prev.Next = curr.Next
				curr = prev.Next
				continue
			}
		}
		prev = curr
		curr = curr.Next
	}
}
