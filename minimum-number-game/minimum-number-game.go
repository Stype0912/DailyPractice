package minimum_number_game

import (
	"slices"
)

func numberGame(nums []int) (arr []int) {
	slices.Sort(nums)
	n := len(nums)
	for i := 0; i < n; i += 2 {
		a, b := nums[i], nums[i+1]
		arr = append(arr, b)
		arr = append(arr, a)
	}
	return
}
