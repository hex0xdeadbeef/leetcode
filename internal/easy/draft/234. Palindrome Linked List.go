package easy

import "fmt"

// https://leetcode.com/problems/palindrome-linked-list/description/

/*
	l1 := easy.ListNode{Val: 1}
	l2 := easy.ListNode{Val: 0}
	l3 := easy.ListNode{Val: 0}
	l4 := easy.ListNode{Val: 1}

	l1.Next, l2.Next, l3.Next = &l2, &l3, &l4

	fmt.Println(easy.IsPalindrome(&l1))
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func IsPalindrome(head *ListNode) bool {

	var (
		count = 1
		// originalHead = head
		bucket int
	)

	for head.Next != nil {
		count++
		bucket = head.Val
		head.Val = head.Next.Val
		head.Next.Val = bucket

		head = head.Next
	}

	fmt.Println(count)

	return true
}

func even() {

}

func odd() {

}
