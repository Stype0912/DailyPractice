package zai_pai_xu_shu_zu_zhong_cha_zhao_shu_zi_lcof

func search(nums []int, target int) int {
	return searchTarget(nums, target, 0)
}

func searchTarget(nums []int, target int, ans int) int {
	if len(nums) == 0 {
		return ans
	}
	if len(nums) == 1 && nums[0] == target {
		return ans + 1
	}
	if target < nums[0] || target > nums[len(nums)-1] {
		return ans
	}
	left := nums[:len(nums)/2]
	right := nums[len(nums)/2:]
	ans = searchTarget(left, target, ans)
	ans = searchTarget(right, target, ans)
	return ans
}
