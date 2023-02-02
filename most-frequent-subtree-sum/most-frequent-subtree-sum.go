package most_frequent_subtree_sum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findFrequentTreeSum(root *TreeNode) (res []int) {
	maxFreq := 0
	vals := map[int]int{}
	var dfs func(node *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		sum := root.Val + dfs(root.Left) + dfs(root.Right)
		vals[sum]++
		if vals[sum] > maxFreq {
			maxFreq = vals[sum]
		}
		return sum
	}
	dfs(root)
	for key, value := range vals {
		if value == maxFreq {
			res = append(res, key)
		}
	}
	return
}
