package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
	len  int
}

func (l *LinkedList) prepend(node *Node) {
	second := l.head
	l.head = node
	l.head.next = second

	l.len++
}

func (l *LinkedList) print() {
	if l.len == 0 {
		return
	}

	node := l.head
	length := l.len
	for length > 0 {
		fmt.Printf("%d ", node.data)
		node = node.next
		length--
	}
	fmt.Println()
}

func (l *LinkedList) deleteVal(val int) {
	if l.len == 0 {
		return
	}

	if l.head.data == val {
		if l.head.next == nil {
			l.head = nil
			l.len--
			return
		}

		n := l.head.next
		l.head = n
		l.len--
	}

	length := l.len
	node := l.head

	for i := 0; i <= l.len; i++ {
		if node.next != nil && node.next.data == val {
			node.next = node.next.next
			length--
		}

	}
	l.len = length
}

func main() {
	list := LinkedList{}
	node1 := &Node{data: 10}
	node2 := &Node{data: 15}
	node3 := &Node{data: 15}
	node4 := &Node{data: 30}
	node5 := &Node{data: 20}
	list.prepend(node1)
	list.prepend(node2)
	list.prepend(node3)
	list.prepend(node4)
	list.prepend(node5)

	list.print()

	list.deleteVal(15)
	list.print()
}

// 20 15 10 10 12 15 34
