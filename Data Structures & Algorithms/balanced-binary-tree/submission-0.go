/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
    
   balance := true
   
   dfs(root, &balance)

   return balance

    
}

func dfs(current *TreeNode, balanced *bool) int {
	if current == nil{
		return 0
	}

	leftTreeHeight := dfs(current.Left,balanced) 
	rightTreeHeight := dfs(current.Right,balanced)

	if leftTreeHeight - rightTreeHeight > 1 || rightTreeHeight - leftTreeHeight > 1 {
		*balanced = false
	}

	return max(leftTreeHeight,rightTreeHeight) + 1
}

func max(a ,b int) int{
	if a>b {
		return a
	}

	return b
}
