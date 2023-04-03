package _sum_ii

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	sumMap1 := map[int]int{}
	sumMap2 := map[int]int{}
	for _, num1 := range nums1 {
		for _, num2 := range nums2 {
			sumMap1[num1+num2]++
		}
	}
	for _, num3 := range nums3 {
		for _, num4 := range nums4 {
			sumMap2[num3+num4]++
		}
	}
	ans := 0
	for key, value := range sumMap1 {
		ans += value * sumMap2[-1*(key)]
	}
	return ans
}
