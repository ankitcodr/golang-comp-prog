package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	result := head
	filtered := head
	for ; head != nil; head = head.Next {
		if head.Next != nil {
			if filtered.Val == head.Next.Val {
				continue
			}
		}
		filtered.Next = head.Next
		filtered = filtered.Next
	}
	return result
}

func DeleteDuplicatesSolution(head *ListNode) *ListNode {
	return deleteDuplicates(head)
}
