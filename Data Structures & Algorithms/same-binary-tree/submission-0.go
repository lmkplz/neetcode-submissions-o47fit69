/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSameTree(p *TreeNode, q *TreeNode) bool {
    var treeToArr func([]int, *TreeNode) []int

	treeToArr = func(result []int, node *TreeNode) []int {
		if node == nil {
			result = append(result, 111)	// 111 means nil 
			return result
		}

		result = append(result, node.Val)

		result = treeToArr(result, node.Left)
		result = treeToArr(result, node.Right)

		return result
	}

	pArr := treeToArr([]int{}, p)
	qArr := treeToArr([]int{}, q)


		fmt.Println(pArr)
		
		fmt.Println(qArr)

	if len(pArr) != len(qArr) {
		return false
	}

	for key, val := range pArr {
		if val != qArr[key] {
			return false
		}
	}

	return true
}
