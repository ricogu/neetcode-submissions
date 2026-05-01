/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
    m := make(map[*ListNode]struct{})

    p := head
    for p!=nil{
        if _,found := m[p]; found {
            return true
        }

        m[p] = struct{}{}
        p = p.Next

    }

    return false
    
}
