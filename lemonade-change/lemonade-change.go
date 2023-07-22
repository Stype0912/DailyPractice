package lemonade_change

func lemonadeChange(bills []int) bool {
	cashMap := map[int]int{}
	for _, bill := range bills {
		switch bill {
		case 5:
			cashMap[5]++
		case 10:
			if cashMap[5] > 0 {
				cashMap[10]++
				cashMap[5]--
			} else {
				return false
			}
		case 20:
			if cashMap[5] <= 0 || (cashMap[10] <= 0 && cashMap[5] <= 2) {
				return false
			}
			if cashMap[5] > 0 && cashMap[10] > 0 {
				cashMap[5]--
				cashMap[10]--
				continue
			}
			if cashMap[10] <= 0 && cashMap[5] >= 3 {
				cashMap[5] -= 3
				continue
			}
		}
	}
	return true
}
