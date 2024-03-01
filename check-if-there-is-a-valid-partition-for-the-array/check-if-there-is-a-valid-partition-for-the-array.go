package check_if_there_is_a_valid_partition_for_the_array

func validPartition(nums []int) bool {
	n := len(nums)
	dp := make([]bool, n+1)
	dp[0] = true
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-2] && (nums[i-1] == nums[i-2])
		if i >= 3 {
			dp[i] = dp[i] || (dp[i-3] && (((nums[i-3] == nums[i-1]) && (nums[i-1] == nums[i-2])) || ((nums[i-1] == nums[i-2]+1) && (nums[i-2] == nums[i-3]+1))))
		}
	}
	return dp[n]
}
