package minimum_recolors_to_get_k_consecutive_black_blocks

func minimumRecolors(blocks string, k int) int {
	countMap := map[uint8]int{}
	ans := k
	for i := 0; i <= len(blocks); i++ {
		if i < k-1 {
			countMap[blocks[i]]++
		} else {
			countMap[blocks[i]]++
			if countMap['W'] < ans {
				ans = countMap['W']
			}
			countMap[blocks[i-k+1]]--
		}
	}
	return ans
}
