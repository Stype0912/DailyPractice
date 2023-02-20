package best_poker_hand

func bestHand(ranks []int, suits []byte) string {
	ranksMap := map[int]int{}
	suitsMap := map[byte]int{}
	for _, rank := range ranks {
		ranksMap[rank]++
	}
	for _, suit := range suits {
		suitsMap[suit]++
	}
	if len(suitsMap) == 1 {
		return "Flush"
	}
	if len(ranksMap) == 5 {
		return "High Card"
	}
	for i := range ranksMap {
		if ranksMap[i] > 2 {
			return "Three of a Kind"
		}
	}
	return "Pair"
}
