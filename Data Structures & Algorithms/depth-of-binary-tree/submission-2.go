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

    queue := []*TreeNode{root}
    depth := 0

    for len(queue) > 0 {
        n := len(queue)
        fmt.Println(n)
        for i := 0; i < n; i++ {
            node := queue[i]

            if node.Right != nil { queue = append(queue, node.Right) }
            if node.Left != nil { queue = append(queue, node.Left) }

            fmt.Println(queue, len(queue))
        }
        queue = queue[n:]
        
        depth++
    }
    
    return depth

}
