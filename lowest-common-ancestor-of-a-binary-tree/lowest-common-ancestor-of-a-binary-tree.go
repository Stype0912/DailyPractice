package lowest_common_ancestor_of_a_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var dfs func(*TreeNode, []*TreeNode)
	pAncestors := []*TreeNode{}
	qAncestors := []*TreeNode{}
	dfs = func(node *TreeNode, nodes []*TreeNode) {
		if node == nil {
			return
		}
		if node.Val == p.Val {
			pAncestors = append(nodes, node)
		}
		if node.Val == q.Val {
			qAncestors = append(nodes, node)
		}
		dfs(node.Left, append(nodes, node))
		dfs(node.Right, append(nodes, node))
	}
	dfs(root, []*TreeNode{})
	for i := len(pAncestors) - 1; i >= 0; i-- {
		for j := len(qAncestors) - 1; j >= 0; j-- {
			if pAncestors[i].Val == qAncestors[j].Val {
				return pAncestors[i]
			}
		}
	}
	return nil
}
