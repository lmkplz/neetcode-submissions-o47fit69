/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }

    return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
    
    // stack := []*TreeNode{root}
    // maxDepth := 1
    // depth := 0

    // for len(stack) > 0 {
    //     node := stack[len(stack)-1]
    //     stack = stack[:len(stack)-1]

    //     if node.Right != nil { stack = append(stack, node.Right) }
    //     if node.Left != nil { stack = append(stack, node.Left) }
        
    // }
}
