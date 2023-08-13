package merge_sorted_array

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j := 0, 0
	ans := make([]int, 0, m+n)
	for {
		if i == m {
			ans = append(ans, nums2[j:]...)
			break
		}
		if j == n {
			ans = append(ans, nums1[i:]...)
			break
		}
		if nums1[i] < nums2[j] {
			ans = append(ans, nums1[i])
			i++
		} else {
			ans = append(ans, nums2[j])
			j++
		}
	}
	copy(nums1, ans)
}
