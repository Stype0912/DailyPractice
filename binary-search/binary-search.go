package binary_search

func search(nums []int, target int) int {
	return binarySearch(nums, 0, len(nums)-1, target)
}

func binarySearch(nums []int, start, end, target int) int {
	if start > end {
		return -1
	}
	mid := (start + end) / 2
	if target > nums[mid] {
		return binarySearch(nums, mid+1, end, target)
	} else if target < nums[mid] {
		return binarySearch(nums, start, mid-1, target)
	} else {
		return mid
	}
}
