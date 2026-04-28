/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Result struct {
	balanced bool
	height int
}

func isBalanced(root *TreeNode) bool {

	var dfs func(*TreeNode) Result

	dfs = func(node *TreeNode) Result {
		if node == nil {
			return Result{true, 0}
		}

		left := dfs(node.Left)
		right := dfs(node.Right)

		balanced := left.balanced && right.balanced && abs(left.height-right.height) <= 1

		return Result{balanced, 1 + max(left.height, right.height)}
	}
	
    res := dfs(root)

	return res.balanced
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
