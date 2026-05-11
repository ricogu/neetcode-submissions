/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

 // 0, 1, 2, 3

func reverseList(head *ListNode) *ListNode {
   return reverseNode(head, nil)

}

func reverseNode(current *ListNode, prev *ListNode) *ListNode {
	if current == nil {
		return prev
	}

	next := current.Next
	current.Next = prev

	return reverseNode(next,current)
    

}
