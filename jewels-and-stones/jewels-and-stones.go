package jewels_and_stones

func numJewelsInStones(jewels string, stones string) int {
	jewelMap := map[int32]bool{}
	for _, jewel := range jewels {
		jewelMap[jewel] = true
	}
	ans := 0
	for _, stone := range stones {
		if jewelMap[stone] {
			ans++
		}
	}
	return ans
}
