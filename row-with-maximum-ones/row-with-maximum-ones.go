package row_with_maximum_ones

func rowAndMaximumOnes(mat [][]int) []int {
	ans := []int{0, 0}
	countn := 0
	for i, row := range mat {
		count := 0
		for _, item := range row {
			if item == 1 {
				count++
			}
		}
		if count > countn {
			ans = []int{i, count}
			countn = count
		}
	}
	return ans
}
