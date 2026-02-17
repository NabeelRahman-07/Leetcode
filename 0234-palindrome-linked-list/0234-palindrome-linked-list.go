/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    p1 := head
    p2 := head

    for p2 != nil && p2.Next != nil {
        p1 = p1.Next
        p2 = p2.Next.Next
    }

    if p2 != nil {
        p1 = p1.Next
    }

    reversed := reverse(p1)
    temp := head
    for reversed != nil {
        if temp.Val == reversed.Val {
            temp = temp.Next
            reversed = reversed.Next
        }else{
            return false
        }
    }
    return true 
}

func reverse(head *ListNode) *ListNode {
    var prev *ListNode
    curr := head
    for curr != nil {
        next := curr.Next
        curr.Next = prev
        prev = curr
        curr = next
    }
    return prev
}