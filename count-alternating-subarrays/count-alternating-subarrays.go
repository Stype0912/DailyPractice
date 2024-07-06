package count_alternating_subarrays

func countAlternatingSubarrays(nums []int) int64 {
	n := len(nums)
	dp := make([]int, n)
	dpLen := make([]int, n)
	for i, _ := range nums {
		dp[i] = 1
		dpLen[i] = 1
		if i >= 1 {
			dp[i] = func() int {
				if nums[i] == nums[i-1] {
					dpLen[i] = 1
					return dp[i-1] + 1
				} else {
					dpLen[i] = dpLen[i-1] + 1
					return dp[i-1] + dpLen[i]
				}
			}()
		}
	}
	return int64(dp[n-1])
}
