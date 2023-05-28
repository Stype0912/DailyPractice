package maximum_width_of_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Pair struct {
	node *TreeNode
	pos  int
}

func widthOfBinaryTree(root *TreeNode) int {
	pairs := []Pair{{
		node: root,
		pos:  1,
	}}
	ans := 1
	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	for len(pairs) != 0 {
		ans = max(ans, pairs[len(pairs)-1].pos-pairs[0].pos+1)
		tmp := make([]Pair, len(pairs))
		copy(tmp, pairs)
		//tmp := pairs
		pairs = nil
		for _, pair := range tmp {
			if pair.node.Left != nil {
				pairs = append(pairs, Pair{pair.node.Left, 2 * pair.pos})
			}
			if pair.node.Right != nil {
				pairs = append(pairs, Pair{pair.node.Right, 2*pair.pos + 1})
			}
		}
	}
	return ans
}
