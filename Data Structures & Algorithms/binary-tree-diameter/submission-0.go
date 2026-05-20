/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */



func diameterOfBinaryTree(root *TreeNode) int {

    maxDiameter := 0

    var dfs func(node *TreeNode) int

    dfs = func(node *TreeNode) int {
        if node == nil {
            return 0
        }

        leftHeight := dfs(node.Left)
        rightHeight := dfs(node.Right)

        maxDiameter = max(maxDiameter, leftHeight + rightHeight)

        return max(leftHeight, rightHeight) + 1

    }

    dfs(root)

    return maxDiameter
   
    
}



func max(a, b int) int{
    if a>b {
        return a
    }

    return b
}
