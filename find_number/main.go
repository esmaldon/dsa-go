package main

import (
	"fmt"
)

func NumInList(a []int, number int) bool {
	if len(a) == 0 {
		fmt.Println("Incorrect array input is nil or empty")
		return false
	}
	for _, i := range a {
		if i == number {
			return true
		}
	}
	return false
}

func main() {
	fmt.Printf("A = {1,2,3,4,5}, Number = 5. Is number in list? %v\n", NumInList([]int{1, 2, 3, 4, 5}, 5))
	fmt.Printf("A = {3,3,3,3,3}, Number = 5. Is number in list? %v\n", NumInList([]int{3, 3, 3, 3, 3}, 5))
	fmt.Printf("A = {3,5,3,5,3}, Number = 5. Is number in list? %v\n", NumInList([]int{3, 5, 3, 5, 3}, 5))
	fmt.Printf("A = {4,2,22,-10,3}, Number = 5. Is number in list? %v\n", NumInList([]int{4, 2, 22, -10, 3}, -10))
	fmt.Printf("A = nil, number = 1.  Is number in list? %v\n", NumInList(nil, 1))
}
