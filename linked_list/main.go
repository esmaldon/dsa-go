package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head *node
}

func newLinkedList() linkedList {
	l := linkedList{}
	return l
}

func (n *node) hasNext() bool {
	return n.next != nil
}

func (n *linkedList) addData(d int) {

	// if liked list is empty then start with first node
	if n.head == nil {
		fmt.Printf("Starting linked list with node value %v\n", d)
		n.head = &node{
			data: d,
		}
		return
	}

	p := n.head

	for p.next != nil {
		fmt.Printf("Moving to next node from %v\n", p.data)
		p = p.next
	}

	fmt.Printf("Creating Node %v\n", d)
	p.next = &node{
		data: d,
	}
}

func (n *linkedList) removeData(d int) {
	if n.head == nil {
		return
	}

	//head verification as there is no previous node
	// 1. if head node is the one to delete but hast next node then move the head
	// 2. if head node is the one to delete but doesn't have next node then head pointer with nil
	if n.head.data == d {
		if n.head.next != nil {
			n.head = n.head.next
			return
		}
		n.head = nil
		return
	}

	// verification through linked list
	// having a pointer on previous node and another pointer that points to the next node that will be checked as fast moving pointer
	previous := n.head
	fast := n.head.next.next

	for previous.next != nil {
		if previous.next.data == d {
			if fast != nil {
				previous.next = fast
				return
			}
			previous.next = nil
		}
		previous = previous.next
		fast = fast.next
	}

}

func (l *linkedList) printList() {
	if l.head == nil {
		fmt.Print("linked list is empty")
	}
	p := l.head

	for p.next != nil {
		fmt.Printf("%v", p.data)
		p = p.next
	}
	fmt.Printf("%v\n", p.data)
}

func (l *linkedList) removeDuplicate() {
	visited := make(map[int]bool)

	previous := l.head
	fast := l.head.next

	visited[previous.data] = true

	for fast.next != nil {
		if visited[fast.data] {
			if fast.next != nil {
				previous.next = fast.next
				fast = fast.next
				continue
			}
			previous.next = nil
			fast.next = previous.next
			continue
		}
		visited[fast.data] = true

		previous = previous.next
		fast = fast.next
	}
	if visited[fast.data] {
		previous.next = nil
	}

}

func main() {

	list := newLinkedList()
	list.addData(1)
	list.addData(2)
	list.addData(3)
	list.addData(2)
	list.addData(3)
	list.printList()
	fmt.Println("Now delete duplicate values")
	list.removeDuplicate()
	list.printList()
}
