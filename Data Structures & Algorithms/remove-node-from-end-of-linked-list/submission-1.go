/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {

    if head.Next == nil && n == 1 {
        return nil
    }

    original := head
    half := []*ListNode{}
    total := 0

    slow := head
    fast := head.Next
    half = append(half, slow)

    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
        half = append(half, slow)
    }

    total = len(half) * 2
    
    if fast == nil {
        total--
    }

    if total-n-1 < 0 {
        return original.Next
    }

    if (total - n) < len(half) {
        half[total-n-1].Next = half[total-n].Next
        return original
    }

    idx := total - n

    for idx > len(half) - 1 {
        slow = slow.Next
        half = append(half, slow)
    }

    half[len(half) - 2].Next = half[len(half)-1].Next
    return original
}
