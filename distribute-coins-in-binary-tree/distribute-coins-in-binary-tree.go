package distribute_coins_in_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func distributeCoins(root *TreeNode) int {
	ans := 0
	var dfs func(*TreeNode) int
	abs := func(i int) int {
		if i > 0 {
			return i
		}
		return -i
	}
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		// moveLeft, moveRight := 0, 0
		moveLeft := dfs(node.Left)
		moveRight := dfs(node.Right)
		ans += abs(moveLeft) + abs(moveRight)
		return moveLeft + moveRight + node.Val - 1
	}
	dfs(root)
	return ans
}
