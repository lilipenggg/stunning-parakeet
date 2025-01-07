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

	l1Cursor := l1
	l2Cursor := l2

	hasNext := true
	incrementOne := false

	for hasNext {
		prev := cursor

		// compute value for current cursor
		value := 0
		if l1Cursor != nil {
			value = l1Cursor.Val
		}
		if l2Cursor != nil {
			value = value + l2Cursor.Val
		}

		if incrementOne {
			value += 1
			incrementOne = false
		}

		if value >= 10 {
			value = int(math.Mod(float64(value), 10))
			incrementOne = true
		}

		cursor = &ListNode{
			Val:  value,
			Next: nil,
		}

		if prev != nil {
			prev.Next = cursor
		} else {
			head = cursor
		}

		hasNext = (l1Cursor != nil && l1Cursor.Next != nil) || (l2Cursor != nil && l2Cursor.Next != nil) || incrementOne

		if hasNext {
			// move the two cursors if there are still more nodes in the list
			if l1Cursor != nil {
				l1Cursor = l1Cursor.Next
			}

			if l2Cursor != nil {
				l2Cursor = l2Cursor.Next
			}
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
