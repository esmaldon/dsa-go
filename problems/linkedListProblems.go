package problems

import (
	"dsa/structures"
	"fmt"
	"strconv"
)

// Sum of Lists: You have a number represented by linked list, each node contains one digit.
// The digits are stored in reverse order
func SumReversedLists(list1, list2 *structures.LinkedList) int {
	head := list1.GetHead()
	var firstNumber, secondNumber string
	for head.Next != nil {
		firstNumber = fmt.Sprintf("%d%s", head.Data, firstNumber)
		head = head.Next
	}
	firstNumber = fmt.Sprintf("%d%s", head.Data, firstNumber)
	fmt.Printf("The first number is %s\n", firstNumber)

	head = list2.GetHead()
	for head.Next != nil {
		secondNumber = fmt.Sprintf("%d%s", head.Data, secondNumber)
		head = head.Next
	}
	secondNumber = fmt.Sprintf("%d%s", head.Data, secondNumber)
	fmt.Printf("The second number is %s\n", secondNumber)

	firstNum, err := strconv.Atoi(firstNumber)
	if err != nil {
		fmt.Println("Could not parse first number string to int")
	}
	secondNum, err := strconv.Atoi(secondNumber)
	if err != nil {
		fmt.Println("Could not parse second number string to int")
	}
	return firstNum + secondNum
}

func SumLists(list1, list2 *structures.LinkedList) int {
	head := list1.GetHead()
	var firstNumber, secondNumber string
	for head.Next != nil {
		firstNumber = fmt.Sprintf("%s%d", firstNumber, head.Data)
		head = head.Next
	}
	firstNumber = fmt.Sprintf("%s%d", firstNumber, head.Data)
	fmt.Printf("The first number is %s\n", firstNumber)

	head = list2.GetHead()
	for head.Next != nil {
		secondNumber = fmt.Sprintf("%s%d", secondNumber, head.Data)
		head = head.Next
	}
	secondNumber = fmt.Sprintf("%s%d", secondNumber, head.Data)
	fmt.Printf("The second number is %s\n", secondNumber)

	firstNum, err := strconv.Atoi(firstNumber)
	if err != nil {
		fmt.Println("Could not parse first number string to int")
	}
	secondNum, err := strconv.Atoi(secondNumber)
	if err != nil {
		fmt.Println("Could not parse second number string to int")
	}
	return firstNum + secondNum
}
