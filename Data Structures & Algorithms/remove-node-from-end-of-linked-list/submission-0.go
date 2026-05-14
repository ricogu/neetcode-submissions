/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	length := 0

	node := head

	for node!=nil {
		node = node.Next
		length++
	}

	indexToRemove := length - n

	current := head
	var prev *ListNode
	for index := 0; index < indexToRemove; index++ {
		prev = current
		current = current.Next
	}

    if prev == nil {
		head = current.Next
	} else {
		prev.Next = current.Next
	}

	return head
    
}
