package find_elements_in_a_contaminated_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type FindElements struct {
	root    *TreeNode
	numsMap map[int]bool
}

func Constructor(root *TreeNode) FindElements {
	root.Val = 0
	var numsMap map[int]bool
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		numsMap[node.Val] = true
		if node.Left != nil {
			node.Left.Val = 2*node.Val + 1
			dfs(node.Left)
		}
		if node.Right != nil {
			node.Right.Val = 2*node.Val + 2
			dfs(node.Right)
		}
	}
	dfs(root)
	return FindElements{
		root:    root,
		numsMap: numsMap,
	}
}

func (this *FindElements) Find(target int) bool {
	return this.numsMap[target]
}

/**
 * Your FindElements object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Find(target);
 */
