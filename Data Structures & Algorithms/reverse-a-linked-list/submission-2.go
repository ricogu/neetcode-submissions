/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//0 -> 1 -> 2 -> 3
//0 <- 1 <- 2 <- 3


// 3 -> 2 -> 1 -> 0          


func reverseList(head *ListNode) *ListNode {
    var prev *ListNode
    current := head

    for current != nil {
        nxt := current.Next
        current.Next = prev
        prev = current
        current = nxt
    }

    return prev
}
