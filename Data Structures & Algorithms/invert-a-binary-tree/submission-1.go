/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func invertTree(root *TreeNode) *TreeNode {
    if root == nil {
        return root
    }

    leftTree := invertTree(root.Left)
    rightTree := invertTree(root.Right)
    root.Right = leftTree
    root.Left = rightTree

    return root
    
}
