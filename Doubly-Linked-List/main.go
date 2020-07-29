package main

type List struct {
	head, tail *Node
}

type Node struct {
	value      int
	next, prev *Node
}

func (l *List) Push(value int) {
	node := &Node{value: value}
	if l.head == nil {
		l.head = node
	} else {
		l.tail.next = node
		node.prev = l.tail
	}

	l.tail = node
}

func main() {
	l := &List{}
	l.Push(1)
	l.Push(2)
	l.Push(3)
	l.Push(4)
}

// type List struct {
// 	head *Node
// 	tail *Node
// }

// type Node struct {
// 	value int
// 	next  *Node
// 	prev  *Node
// }

// func (l *List) Push(val int) {
// 	node := &Node{value: val}
// 	if l.head == nil {
// 		l.head = node
// 	} else {
// 		l.tail.next = node
// 		node.prev = l.tail
// 	}
// 	l.tail = node
// }

// func (l *List) First() *Node {
// 	return l.head
// }

// func (l *List) Last() *Node {
// 	return l.tail
// }

// func (n *Node) Next() *Node {
// 	return n.next
// }

// func (n *Node) Prev() *Node {
// 	return n.prev
// }

// func main() {
// 	l := &List{}
// 	for i := 1; i <= 5; i++ {
// 		l.Push(i)
// 	}
// 	n := l.First()
// 	for {
// 		println(n.value)
// 		n = n.Next()
// 		if n == nil {
// 			break
// 		}
// 	}

// 	n = l.Last()
// 	for {
// 		println(n.value)
// 		n = n.Prev()
// 		if n == nil {
// 			break
// 		}
// 	}

// }
