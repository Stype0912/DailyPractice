package main

func findLengthOfShortestSubarray(arr []int) int {
	length := len(arr)
	if length == 1 {
		return 0
	}
	left, right := -1, -1
	for i := 1; i < length; i++ {
		if i == length-1 {
			return 0
		}
		if arr[i] < arr[i-1] {
			left = i
			break
		}
	}
	for j := length - 1; j >= 1; j-- {
		if arr[j-1] > arr[j] {
			right = j
			break
		}
	}
	ans := right - left
	i := 0
	j := right
	for {
		if i == left || j == length {
			break
		}
		if arr[i] > arr[j] {
			ans++
			j++
		}
		i++
	}
	return ans
}