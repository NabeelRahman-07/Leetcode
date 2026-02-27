/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return nil
    }
    curr := head
    seen := map[*ListNode]bool{}
    for curr != nil {
        if _, ok := seen[curr]; ok {
            return curr
        }
        seen[curr] = true
        curr = curr.Next
    }
    return nil
}