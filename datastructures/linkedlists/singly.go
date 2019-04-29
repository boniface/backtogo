package linkedlists

type ListNode struct {
	next *ListNode
	data interface{}
}

type LinkedList struct {
	head *ListNode
	size int
}
