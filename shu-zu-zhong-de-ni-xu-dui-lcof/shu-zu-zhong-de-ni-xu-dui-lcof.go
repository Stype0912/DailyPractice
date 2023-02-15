package shu_zu_zhong_de_ni_xu_dui_lcof

func reversePairs(nums []int) int {
	res := 0
	var mergeSort func(int, int) int
	mergeSort = func(left, right int) int {
		if left >= right {
			return 0
		}
		m := (left + right) / 2
		res = mergeSort(left, m) + mergeSort(m+1, right)
		i, j := left, m+1
		tmp := []int{}
		for k := left; k <= right; k++ {
			if i == m+1 {
				tmp = append(tmp, nums[j])
				j++
				continue
			} else if j == right+1 {
				tmp = append(tmp, nums[i])
				i++
				continue
			}
			if nums[i] > nums[j] {
				tmp = append(tmp, nums[j])
				res += m - i + 1
				j++
			} else {
				tmp = append(tmp, nums[i])
				i++
			}
		}
		for k := left; k <= right; k++ {
			nums[k] = tmp[k-left]
		}
		return res
	}
	return mergeSort(0, len(nums)-1)
}
