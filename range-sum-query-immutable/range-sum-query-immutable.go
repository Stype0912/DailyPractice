package range_sum_query_immutable

type NumArray struct {
	prefixSum []int
}

func Constructor(nums []int) NumArray {
	n := len(nums)
	na := NumArray{
		prefixSum: make([]int, n),
	}
	if n == 0 {
		return na
	}
	na.prefixSum[0] = nums[0]
	for i := 1; i < n; i++ {
		na.prefixSum[i] = nums[i] + na.prefixSum[i-1]
	}
	return na
}

func (this *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return this.prefixSum[right]
	}
	return this.prefixSum[right] - this.prefixSum[left-1]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
