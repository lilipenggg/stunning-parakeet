package main

import (
	"fmt"
	"math"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var cursor *ListNode
	var head *ListNode

	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		prev := cursor

		// compute value for current cursor
		value := 0
		if l1 != nil {
			value = l1.Val
		}
		if l2 != nil {
			value = value + l2.Val
		}
		value += carry

		// reset carry and value if value is larger than 10
		carry = int(math.Floor(float64(value / 10)))
		value = int(math.Mod(float64(value), 10))

		cursor = &ListNode{
			Val:  value,
			Next: nil,
		}

		if prev != nil {
			prev.Next = cursor
		} else {
			// record the head of the linked list
			head = cursor
		}

		// move the cursors to next ones
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	return head
}

func printNode(head *ListNode) {
	var res []int

	next := head
	for next != nil {
		res = append(res, next.Val)
		next = next.Next
	}

	fmt.Println(res)
}

func buildListNode(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	var head *ListNode

	var cursor *ListNode
	for _, n := range nums {
		prev := cursor

		cursor = &ListNode{
			Val:  n,
			Next: nil,
		}

		if prev != nil {
			prev.Next = cursor
		} else {
			head = cursor
		}
	}

	return head
}

func main() {
	l1 := buildListNode([]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1})
	l2 := buildListNode([]int{5, 6, 4})

	result := addTwoNumbers(l1, l2)

	printNode(result)

	l3 := buildListNode([]int{2, 4, 3})
	l4 := buildListNode([]int{5, 6, 4})

	result = addTwoNumbers(l3, l4)

	printNode(result)

	l5 := buildListNode([]int{9, 9, 9, 9, 9, 9, 9})
	l6 := buildListNode([]int{9, 9, 9, 9})

	result = addTwoNumbers(l5, l6)

	printNode(result)

	l7 := buildListNode([]int{0})
	l8 := buildListNode([]int{0})

	result = addTwoNumbers(l7, l8)

	printNode(result)
}
