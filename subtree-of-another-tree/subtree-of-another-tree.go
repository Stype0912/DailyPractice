package subtree_of_another_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	var dfs func(*TreeNode) bool
	dfs = func(node *TreeNode) bool {
		if node == nil {
			return false
		}
		return check(node, subRoot) || dfs(node.Left) || dfs(node.Right)
	}
	return dfs(root)
}

func check(node1, node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}
	if node1 == nil && node2 != nil {
		return false
	}
	if node1 != nil && node2 == nil {
		return false
	}
	if node1.Val != node2.Val {
		return false
	}
	return check(node1.Left, node2.Left) && check(node1.Right, node2.Right)
}
