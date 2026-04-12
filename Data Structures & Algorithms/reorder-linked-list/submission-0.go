/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
    queue := []*ListNode{}
    stack := []*ListNode{}

    original := head
    next := head

    for next != nil && next.Next != nil {
        next = next.Next.Next
        queue = append(queue, head)
        head = head.Next
        fmt.Println(next)
    }

    fmt.Println(len(queue), queue)

    for head != nil {
        stack = append(stack, head)
        head = head.Next
    }
    fmt.Println(len(stack), stack)

    for i := 0; i < len(queue); i++ {
        tmp1 := queue[i].Next

        queue[i].Next = stack[len(stack)-1-i]
        stack[len(stack)-1-i].Next = tmp1
    }

    stack[0].Next = nil

    head = original
}
