package arithmetic_subarray

import "sort"

type intList []int

func (m intList) Len() int {
	return len(m)
}

func (m intList) Less(i, j int) bool {
	return m[i] < m[j]
}

func (m intList) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	ans := []bool{}
	for i := 0; i < len(l); i++ {
		numsTmp := make([]int, r[i]-l[i]+1)
		copy(numsTmp, nums[l[i]:r[i]+1])
		numsArray := intList(numsTmp)
		sort.Sort(numsArray)
		ans = append(ans, isEqualMinus(numsArray))
	}
	return ans
}

func isEqualMinus(numArray intList) bool {
	if len(numArray) <= 2 {
		return true
	}
	minus := numArray[1] - numArray[0]
	for i := 2; i < len(numArray); i++ {
		if numArray[i]-numArray[i-1] != minus {
			return false
		}
	}
	return true
}
