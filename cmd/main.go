package main

import (
	"dsa/structures"
	"fmt"
)

func main() {

	list := structures.NewLinkedList()
	list.AddData(1)
	list.AddData(2)
	list.AddData(3)
	list.AddData(2)
	list.AddData(3)
	list.PrintList()
	fmt.Println("Now delete duplicate values")
	list.RemoveDuplicate()
	list.PrintList()
	list.AddData(4)
	list.AddData(5)
	list.PrintList()
	fmt.Printf("The kth element to last of %v is %v", 2, list.KthToLast(2))
}
