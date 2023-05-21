package maximum_sum_bst_in_binary_tree

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const MIN = math.MinInt
const MAX = math.MaxInt

type subTree struct {
	isBst  bool
	min    int
	max    int
	subSum int
}

func maxSumBST(root *TreeNode) int {
	ans := 0
	min := func(i, j int) int {
		if i < j {
			return i
		}
		return j
	}
	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	var dfs func(*TreeNode) *subTree
	dfs = func(root *TreeNode) *subTree {
		if root == nil {
			return &subTree{
				isBst:  true,
				min:    MAX,
				max:    MIN,
				subSum: 0,
			}
		}
		leftTree := dfs(root.Left)
		rightTree := dfs(root.Right)
		if leftTree.isBst && rightTree.isBst && leftTree.max < root.Val && rightTree.min > root.Val {
			ans = max(ans, root.Val+leftTree.subSum+rightTree.subSum)
			return &subTree{
				isBst:  true,
				min:    min(root.Val, leftTree.min),
				max:    max(root.Val, rightTree.max),
				subSum: root.Val + leftTree.subSum + rightTree.subSum,
			}
		}
		return &subTree{
			isBst:  false,
			min:    0,
			max:    0,
			subSum: 0,
		}
	}

	dfs(root)
	return ans
}
