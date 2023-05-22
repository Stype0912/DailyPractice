package insufficient_nodes_in_root_to_leaf_paths

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sufficientSubset(root *TreeNode, limit int) *TreeNode {
	if root == nil {
		return nil
	}
	var dfs func(*TreeNode, int) bool
	dfs = func(node *TreeNode, sum int) bool {
		if node == nil {
			return false
		}
		sum += node.Val
		if node.Left == nil && node.Right == nil {
			return sum >= limit
		}
		bool1 := dfs(node.Left, sum)
		bool2 := dfs(node.Right, sum)
		if !bool1 {
			node.Left = nil
		}
		if !bool2 {
			node.Right = nil
		}
		if bool1 || bool2 {
			return true
		} else {
			return false
		}
	}
	if dfs(root, 0) {
		return root
	} else {
		return nil
	}
}
