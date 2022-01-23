package main

func searchRange(nums []int, target int) (res []int) {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	i, j := 0, len(nums)-1
	res = []int{-1, -1}
	flag1, flag2 := true, true
	for {
		if !flag1 && !flag2 {
			return
		}
		if i > j {
			return
		}
		if nums[i] == target && flag1 {
			res[0] = i
			flag1 = false
		}
		if nums[j] == target && flag2 {
			res[1] = j
			flag2 = false
		}
		if flag1 {
			i++
		}
		if flag2 {
			j--
		}
	}
}
