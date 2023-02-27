package decrease_elements_to_make_array_zigzag

func movesToMakeZigzag(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	res1, res2 := 0, 0
	cache1 := make([]int, len(nums))
	cache2 := make([]int, len(nums))
	copy(cache1, nums)
	copy(cache2, nums)
	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			if i == 0 {
				if cache1[i] <= cache1[i+1] {
					res1 += cache1[i+1] - cache1[i] + 1
					cache1[i+1] -= cache1[i+1] - cache1[i] + 1
				}
			} else if i == len(nums)-1 {
				if cache1[i] <= cache1[i-1] {
					res1 += cache1[i-1] - cache1[i] + 1
					cache1[i-1] -= cache1[i-1] - cache1[i] + 1
				}
			} else {
				if cache1[i] <= cache1[i-1] {
					res1 += cache1[i-1] - cache1[i] + 1
					cache1[i-1] -= cache1[i-1] - cache1[i] + 1
				}
				if cache1[i] <= cache1[i+1] {
					res1 += cache1[i+1] - cache1[i] + 1
					cache1[i+1] -= cache1[i+1] - cache1[i] + 1
				}
			}
		} else {
			if i == len(nums)-1 {
				if cache2[i] <= cache2[i-1] {
					res2 += cache2[i-1] - cache2[i] + 1
					cache2[i-1] -= cache2[i-1] - cache2[i] + 1
				}
			} else {
				if cache2[i] <= cache2[i-1] {
					res2 += cache2[i-1] - cache2[i] + 1
					cache2[i-1] -= cache2[i-1] - cache2[i] + 1
				}
				if cache2[i] <= cache2[i+1] {
					res2 += cache2[i+1] - cache2[i] + 1
					cache2[i+1] -= cache2[i+1] - cache2[i] + 1
				}
			}
		}
	}
	return min(res1, res2)
}

func min(i, j int) int {
	if i < j {
		return i
	} else {
		return j
	}
}
