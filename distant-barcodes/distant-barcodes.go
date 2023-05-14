package distant_barcodes

import "sort"

func rearrangeBarcodes(barcodes []int) []int {
	codeCntMap := map[int]int{}
	for _, code := range barcodes {
		codeCntMap[code]++
	}
	sort.Slice(barcodes, func(i, j int) bool {
		if codeCntMap[barcodes[i]] == codeCntMap[barcodes[j]] {
			return barcodes[i] > barcodes[j]
		}
		return codeCntMap[barcodes[i]] > codeCntMap[barcodes[j]]
	})
	ans := make([]int, len(barcodes))
	j := 0
	for i := 0; i < len(barcodes); i += 2 {
		ans[i] = barcodes[j]
		j++
	}
	for i := 1; i < len(barcodes); i += 2 {
		ans[i] = barcodes[j]
		j++
	}
	return ans
}
