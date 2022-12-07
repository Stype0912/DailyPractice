package maximum_product_of_three_numbers

import "sort"

type Nums []int

func (nums Nums) Len() int {
	return len(nums)
}

func (nums Nums) Swap(i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

func (nums Nums) Less(i, j int) bool {
	return nums[i] > nums[j]
}

func maximumProduct(nums []int) int {
	numsNew := Nums(nums)
	sort.Sort(numsNew)
	return max(numsNew[0]*numsNew[1]*numsNew[2], numsNew[0]*numsNew[len(numsNew)-1]*numsNew[len(numsNew)-2])
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
