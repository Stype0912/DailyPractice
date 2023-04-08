package kth_largest_element_in_an_array

import (
	"math/rand"
	"time"
)

func findKthLargest(nums []int, k int) int {
	rand.Seed(time.Now().UnixNano())
	return quickSelect(nums, 0, len(nums)-1, k)
}

func quickSelect(arr []int, l, r, index int) int {
	q := randomPartition(arr, l, r)
	if q == index-1 {
		return arr[q]
	} else if q < index {
		return quickSelect(arr, q+1, r, index)
	} else {
		return quickSelect(arr, l, q-1, index)
	}
}

func randomPartition(arr []int, l, r int) int {
	i := rand.Int()%(r-l+1) + l
	arr[i], arr[r] = arr[r], arr[i]
	return partition(arr, l, r)
}

func partition(arr []int, l, r int) int {
	x := arr[r]
	i := l - 1
	for j := l; j < r; j++ {
		if arr[j] >= x {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[r] = arr[r], arr[i+1]
	return i + 1
}
