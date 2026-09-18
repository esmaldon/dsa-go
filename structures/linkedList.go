package structures

import (
	"fmt"
)

type nodeS struct {
	data int
	next *nodeS
}

type linkedList struct {
	head *nodeS
}

func NewLinkedList() linkedList {
	l := linkedList{}
	return l
}

func (n *nodeS) hasNext() bool {
	return n.next != nil
}

func (n *linkedList) AddData(d int) {

	// if liked list is empty then start with first nodeS
	if n.head == nil {
		fmt.Printf("Starting linked list with nodeS value %v\n", d)
		n.head = &nodeS{
			data: d,
		}
		return
	}

	p := n.head

	for p.next != nil {
		fmt.Printf("Moving to next nodeS from %v\n", p.data)
		p = p.next
	}

	fmt.Printf("Creating nodeS %v\n", d)
	p.next = &nodeS{
		data: d,
	}
}

func (n *linkedList) RemoveData(d int) {
	if n.head == nil {
		return
	}

	//head verification as there is no previous nodeS
	// 1. if head nodeS is the one to delete but hast next nodeS then move the head
	// 2. if head nodeS is the one to delete but doesn't have next nodeS then head pointer with nil
	if n.head.data == d {
		if n.head.next != nil {
			n.head = n.head.next
			return
		}
		n.head = nil
		return
	}

	// verification through linked list
	// having a pointer on previous nodeS and another pointer that points to the next nodeS that will be checked as fast moving pointer
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

func (l *linkedList) PrintList() {
	if l.head == nil {
		fmt.Print("linked list is empty")
	}
	p := l.head

	for p.next != nil {
		fmt.Printf("%v", p.data)
		p = p.next
	}
	fmt.Printf("%v\n", p.data)
	fmt.Printf("size %v\n", l.size())
}

func (l *linkedList) RemoveDuplicate() {
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

func (l *linkedList) size() int {
	counter := 0
	h := l.head

	if h == nil {
		return 0
	}

	for h.next != nil {
		counter = counter + 1
		h = h.next
	}

	counter = counter + 1
	return counter
}

func (l *linkedList) KthToLast(d int) int {
	size := l.size()

	if d > size {
		fmt.Printf("The element %v is grather than size")
		return 0
	}

	k := (size - d) + 1
	h := l.head
	for i := 1; i < k; i++ {
		h = h.next
	}
	return h.data
}

type nodeD struct {
	data           int
	previous, next *nodeD
}

type linkedListDouble struct {
	head *nodeD
}
