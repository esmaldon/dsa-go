package structures

import (
	"fmt"
)

type NodeS struct {
	Data int
	Next *NodeS
}

type LinkedList struct {
	head *NodeS
}

func NewLinkedList() LinkedList {
	l := LinkedList{}
	return l
}

func (n *NodeS) HasNext() bool {
	return n.Next != nil
}

func (l *LinkedList) GetHead() *NodeS {
	return l.head
}

func (n *LinkedList) AddData(d int) {

	// if liked list is empty then start with first nodeS
	if n.head == nil {
		fmt.Printf("Starting linked list with nodeS value %v\n", d)
		n.head = &NodeS{
			Data: d,
		}
		return
	}

	p := n.head

	for p.Next != nil {
		fmt.Printf("Moving to next nodeS from %v\n", p.Data)
		p = p.Next
	}

	fmt.Printf("Creating nodeS %v\n", d)
	p.Next = &NodeS{
		Data: d,
	}
}

func (n *LinkedList) RemoveData(d int) {
	if n.head == nil {
		return
	}

	//head verification as there is no previous nodeS
	// 1. if head nodeS is the one to delete but hast next nodeS then move the head
	// 2. if head nodeS is the one to delete but doesn't have next nodeS then head pointer with nil
	if n.head.Data == d {
		if n.head.Next != nil {
			n.head = n.head.Next
			return
		}
		n.head = nil
		return
	}

	// verification through linked list
	// having a pointer on previous nodeS and another pointer that points to the next nodeS that will be checked as fast moving pointer
	previous := n.head
	fast := n.head.Next.Next

	for previous.Next != nil {
		if previous.Next.Data == d {
			if fast != nil {
				previous.Next = fast
				return
			}
			previous.Next = nil
		}
		previous = previous.Next
		fast = fast.Next
	}

}

func (l *LinkedList) PrintList() {
	if l.head == nil {
		fmt.Print("linked list is empty")
	}
	p := l.head

	for p.Next != nil {
		fmt.Printf("%v", p.Data)
		p = p.Next
	}
	fmt.Printf("%v\n", p.Data)
	fmt.Printf("size %v\n", l.Size())
}

func (l *LinkedList) RemoveDuplicate() {
	visited := make(map[int]bool)

	previous := l.head
	fast := l.head.Next

	visited[previous.Data] = true

	for fast.Next != nil {
		if visited[fast.Data] {
			if fast.Next != nil {
				previous.Next = fast.Next
				fast = fast.Next
				continue
			}
			previous.Next = nil
			fast.Next = previous.Next
			continue
		}
		visited[fast.Data] = true

		previous = previous.Next
		fast = fast.Next
	}
	if visited[fast.Data] {
		previous.Next = nil
	}

}

func (l *LinkedList) Size() int {
	counter := 0
	h := l.head

	if h == nil {
		return 0
	}

	for h.Next != nil {
		counter = counter + 1
		h = h.Next
	}

	counter = counter + 1
	return counter
}

func (l *LinkedList) KthToLast(d int) int {
	size := l.Size()

	if d > size {
		fmt.Printf("The element %v is grather than size")
		return 0
	}

	k := (size - d) + 1
	h := l.head
	for i := 1; i < k; i++ {
		h = h.Next
	}
	return h.Data
}

type nodeD struct {
	Data           int
	previous, Next *nodeD
}

type linkedListDouble struct {
	head *nodeD
}
