package delete_nodes_and_return_forest

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	ans := []*TreeNode{}
	deleteMap := map[int]int{}
	for _, num := range to_delete {
		deleteMap[num]++
	}
	var dfs func(*TreeNode, bool) bool
	dfs = func(node *TreeNode, isRoot bool) bool {
		if node == nil {
			return false
		}
		if deleteMap[node.Val] == 0 && isRoot {
			ans = append(ans, node)
			if dfs(node.Left, false) {
				node.Left = nil
			}
			if dfs(node.Right, false) {
				node.Right = nil
			}
			return false
		} else if deleteMap[node.Val] == 0 && !isRoot {
			if dfs(node.Left, false) {
				node.Left = nil
			}
			if dfs(node.Right, false) {
				node.Right = nil
			}
			return false
		} else if deleteMap[node.Val] != 0 {
			dfs(node.Left, true)
			dfs(node.Right, true)
			return true
		}
		return false
	}
	dfs(root, true)
	return ans
}
