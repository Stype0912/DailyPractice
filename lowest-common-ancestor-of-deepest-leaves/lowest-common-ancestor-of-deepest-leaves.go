package lowest_common_ancestor_of_deepest_leaves

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	levelMap := map[int][]*TreeNode{}
	fatherMap := map[int][]*TreeNode{}
	var dfs func(node *TreeNode, ancestor []*TreeNode, level int)
	maxLevel := 0
	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	dfs = func(node *TreeNode, ancestor []*TreeNode, level int) {
		if node == nil {
			return
		}
		level++
		maxLevel = max(maxLevel, level)
		levelMap[level] = append(levelMap[level], node)
		newAncestor := make([]*TreeNode, len(ancestor)+1)
		for i := 0; i < len(ancestor); i++ {
			newAncestor[i] = ancestor[i]
		}
		newAncestor[len(ancestor)] = node
		fatherMap[node.Val] = newAncestor
		dfs(node.Left, append(ancestor, node), level)
		dfs(node.Right, append(ancestor, node), level)
	}
	dfs(root, []*TreeNode{}, 0)
	deepestNodes := levelMap[maxLevel]
	lowestCommon := func(node1, node2 *TreeNode) *TreeNode {
		ancestor1, ancestor2 := fatherMap[node1.Val], fatherMap[node2.Val]
		for i := len(ancestor1) - 1; i >= 0; i-- {
			for j := len(ancestor2) - 1; j >= 0; j-- {
				if ancestor1[i].Val == ancestor2[j].Val {
					return ancestor1[i]
				}
			}
		}
		return nil
	}
	tmp := deepestNodes[0]
	for i := 1; i < len(deepestNodes); i++ {
		tmp = lowestCommon(tmp, deepestNodes[i])
	}

	return tmp
}
