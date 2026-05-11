/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {

	m := make(map[*ListNode]int)

	curr := head

	for curr != nil {
        if m[curr] == 1{
			return true
		}

		m[curr] = 1
        curr = curr.Next

	}

	return false
    
}
