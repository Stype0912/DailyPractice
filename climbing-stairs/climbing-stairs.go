package main

func climbStairs(n int) int {
	ansList := []int{1, 2}
	if n <= 2 {
		return ansList[n-1]
	}
	for i := 2; i < n; i++ {
		ansList = append(ansList, ansList[i-1]+ansList[i-2])
	}
	return ansList[n-1]
}

