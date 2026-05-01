/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    dummy := &ListNode{}
    tail := dummy

    l1,l2 := list1,list2

    for l1 != nil && l2 != nil{
        if l1.Val < l2.Val {
            tail.Next = l1
            l1 = l1.Next
        } else {
            tail.Next = l2
            l2 = l2.Next
        }

        tail = tail.Next
    }

    if l1 != nil {
        tail.Next = l1
    }

    if l2 != nil {
        tail.Next = l2
    }

    return dummy.Next

}
