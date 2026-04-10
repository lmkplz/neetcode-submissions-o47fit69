/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    var nxt1 *ListNode = list1
	var nxt2 *ListNode = list2
	var list *ListNode = nil
	var cur *ListNode = nil

	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	if list1.Val > list2.Val {
		cur = list2
		nxt2 = list2.Next
	} else {
		cur = list1
		nxt1 = list1.Next
	}

	list = cur

	for nxt1 != nil && nxt2 != nil {
		if nxt1.Val > nxt2.Val {
			cur.Next = nxt2
			cur = nxt2
			nxt2 = nxt2.Next
		} else {
			cur.Next = nxt1
			cur = nxt1
			nxt1 = nxt1.Next
		}
	}

	if nxt1 == nil {
		cur.Next = nxt2
	} else {
		cur.Next = nxt1
	}

	return list
}
