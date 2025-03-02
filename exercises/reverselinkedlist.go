package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || left == right {
		return head
	}
	dummy := &ListNode{Next: head}
	prev := dummy
	for i := 0; i < left-1; i++ {
		prev = prev.Next
	}
	current := prev.Next

	for i := 0; i < right-left; i++ {
		next := current.Next
		current.Next = next.Next
		next.Next = prev.Next
		prev.Next = next
	}
	return dummy.Next
}

//func main() {
//	head := &ListNode{Val: 1}
//	head.Next = &ListNode{Val: 2}
//	head.Next.Next = &ListNode{Val: 3}
//	head.Next.Next.Next = &ListNode{Val: 4}
//	head.Next.Next.Next.Next = &ListNode{Val: 5}
//	head = reverseBetween(head, 2, 4)
//	current := head
//	for current != nil {
//		fmt.Println(current.Val)
//		current = current.Next
//	}
//}
