/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	
	m := make(map[*ListNode]bool)

    for head.Next != nil {
		nxt := head.Next
		if m[nxt] == false {
			m[nxt] = true
		} else {
			return true
		}

		head = head.Next
	}

	return false
}
