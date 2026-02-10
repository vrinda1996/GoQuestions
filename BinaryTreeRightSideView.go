/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
    var res []int
    dfs(root, 0, &res)
    return res
}

func dfs(node *TreeNode, depth int, res *[]int){
    if node == nil{
        return
    }
    if len(*res) == depth {
        *res = append(*res, node.Val)
    }
    if node.Right!=nil {
        dfs(node.Right, depth + 1, res)
    }
    if node.Left!=nil {
        dfs(node.Left, depth + 1, res)
    }
}
