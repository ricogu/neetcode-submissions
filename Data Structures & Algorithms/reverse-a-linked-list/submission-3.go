/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

 // 0, 1, 2, 3

func reverseList(head *ListNode) *ListNode {
   
    var prev *ListNode
	current := head

	for current != nil{
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	return prev
    
}
