package path_sum_iii

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//func pathSum(root *TreeNode, targetSum int) int {
//	if root == nil {
//		return 0
//	}
//	var dfsPathSum func(*TreeNode, int)
//	ans := 0
//	dfsPathSum = func(root *TreeNode, targetSum int) {
//		if root == nil {
//			return
//		}
//		if targetSum-root.Val == 0 {
//			ans++
//		}
//		dfsPathSum(root.Left, targetSum-root.Val)
//		dfsPathSum(root.Right, targetSum-root.Val)
//	}
//	dfsPathSum(root, targetSum)
//	ans += pathSum(root.Left, targetSum)
//	ans += pathSum(root.Right, targetSum)
//	return ans
//}

func pathSum(root *TreeNode, targetSum int) int {
	preSum := map[int]int{0: 1}
	var dfs func(*TreeNode, int)
	ans := 0
	dfs = func(node *TreeNode, curr int) {
		if node == nil {
			return
		}
		curr += node.Val
		ans += preSum[curr-targetSum]
		preSum[curr]++
		dfs(node.Left, curr)
		dfs(node.Right, curr)
		preSum[curr]--
	}
	dfs(root, 0)
	return ans
}
