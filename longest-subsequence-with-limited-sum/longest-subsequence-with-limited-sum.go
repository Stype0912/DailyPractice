package longest_subsequence_with_limited_sum

import "sort"

type intList []int

func (l intList) Len() int {
	return len(l)
}

func (l intList) Less(i, j int) bool {
	return l[i] < l[j]
}

func (l intList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func answerQueries(nums []int, queries []int) []int {
	numsList := intList(nums)
	sort.Sort(numsList)
	sumList := []int{numsList[0]}
	for i := 1; i < len(numsList); i++ {
		sumList = append(sumList, sumList[i-1]+numsList[i])
	}
	ans := []int{}
	for _, query := range queries {
		idx := sort.SearchInts(sumList, query+1)
		ans = append(ans, idx)
	}
	return ans
}
