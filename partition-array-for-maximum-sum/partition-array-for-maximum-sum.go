package partition_array_for_maximum_sum

func maxSumAfterPartitioning(arr []int, k int) int {
	dp := make([]int, len(arr))
	subArrayMax := -1
	for i := 0; i < k; i++ {
		subArrayMax = max(subArrayMax, arr[i])
		dp[i] = subArrayMax * (i + 1)
	}
	for i := k; i < len(arr); i++ {
		subArrayMax = -1
		for j := i - k; j < i; j++ {
			subArrayMax = max(subArrayMax, dp[j]+sumOfSubArray(arr[j+1:i+1]))
		}
		dp[i] += subArrayMax
	}
	return dp[len(arr)-1]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func sumOfSubArray(subArray []int) int {
	subArrayMax := -1
	for i := 0; i < len(subArray); i++ {
		subArrayMax = max(subArrayMax, subArray[i])
	}
	return subArrayMax * len(subArray)
}
