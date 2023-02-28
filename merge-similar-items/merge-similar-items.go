package merge_similar_items

import "sort"

type intType [][]int

func (it intType) Len() int {
	return len(it)
}
func (it intType) Less(i int, j int) bool {
	return it[i][0] < it[j][0]
}

func (it intType) Swap(i int, j int) {
	it[i], it[j] = it[j], it[i]
}

func mergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {
	valueMap := map[int]int{}
	for _, item := range items1 {
		valueMap[item[0]] = item[1]
	}
	for _, item := range items2 {
		if _, ok := valueMap[item[0]]; !ok {
			valueMap[item[0]] = item[1]
		} else {
			valueMap[item[0]] += item[1]
		}
	}
	ans := [][]int{}
	for key, value := range valueMap {
		ans = append(ans, []int{key, value})
	}
	sort.Sort(intType(ans))
	return ans
}
