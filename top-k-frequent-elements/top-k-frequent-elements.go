package top_k_frequent_elements

import "sort"

type NumMap [][2]int

func (n NumMap) Len() int {
	return len(n)
}

func (n NumMap) Less(i, j int) bool {
	return n[i][1] > n[j][1]
}

func (n NumMap) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func topKFrequent(nums []int, k int) []int {
	numMap := map[int]int{}
	for _, num := range nums {
		numMap[num]++
	}
	numCount := [][2]int{}
	for key, value := range numMap {
		numCount = append(numCount, [2]int{key, value})
	}
	sort.Sort(NumMap(numCount))
	ans := []int{}
	for i := 0; i < k; i++ {
		ans = append(ans, numCount[i][0])
	}
	return ans
}
