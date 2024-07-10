package next_permutation

func nextPermutation(nums []int) {
	n := len(nums)
	i, j := n-1, n-1
	reverse := func(list []int) {
		m := len(list)
		for k := 0; k < m/2; k++ {
			list[k], list[m-1-k] = list[m-1-k], list[k]
		}
	}
	for ; i > 0; i-- {
		if nums[i-1] < nums[i] {
			break
		}
	}
	if i == 0 {
		reverse(nums)
		return
	}
	i--
	for ; j > i; j-- {
		if nums[j] > nums[i] {
			nums[i], nums[j] = nums[j], nums[i]
			reverse(nums[i+1:])
			break
		}
	}
}
