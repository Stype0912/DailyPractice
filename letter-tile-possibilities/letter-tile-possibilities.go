package letter_tile_possibilities

func numTilePossibilities(tiles string) int {
	wordMapExist := map[string]int{}
	wordMap := map[byte]int{}
	for i := 0; i < len(tiles); i++ {
		wordMap[tiles[i]]++
	}
	ans := 0
	var dfs func(string, map[byte]int)
	dfs = func(preStr string, m map[byte]int) {
		if preStr != "" && wordMapExist[preStr] == 0 {
			ans++
			wordMapExist[preStr]++
		}
		for key, _ := range m {
			if m[key] > 0 {
				m[key]--
				dfs(preStr+string(key), m)
				m[key]++
			}
		}
	}
	dfs("", wordMap)
	return ans
}
