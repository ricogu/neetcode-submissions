/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//0 -> 1 -> 2 -> 3

// [0,1,2,3]

func reverseList(head *ListNode) *ListNode {
    p := head
    list := []int{}
    for p != nil {
        list = append(list,p.Val)
        p = p.Next
    }

    if len(list)<2{
        return head
    }

    newHead := &ListNode{Val: list[len(list)-1]}
    current := newHead
    for i:= len(list)-2; i>=0; i-- {
        current.Next = &ListNode{Val: list[i]}

        current = current.Next

    }

    return newHead


}
