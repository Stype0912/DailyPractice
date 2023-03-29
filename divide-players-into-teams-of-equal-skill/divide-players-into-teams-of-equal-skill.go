package divide_players_into_teams_of_equal_skill

func dividePlayers(skill []int) int64 {
	skillSum := 0
	skillMap := map[int]int{}
	for _, skillNum := range skill {
		skillSum += skillNum
		skillMap[skillNum]++
	}
	groupNum := len(skill) / 2
	if skillSum%groupNum != 0 {
		return -1
	}
	groupSkill := skillSum / groupNum
	ansGroup := [][2]int{}
	ans := int64(0)
	for i := 0; i < len(skill); i++ {
		if skillMap[skill[i]] == 0 {
			continue
		} else {
			skillMap[skill[i]]--
			if skillMap[groupSkill-skill[i]] <= 0 {
				return -1
			} else {
				skillMap[groupSkill-skill[i]]--
				ansGroup = append(ansGroup, [2]int{skill[i], groupSkill - skill[i]})
				ans += int64(skill[i] * (groupSkill - skill[i]))
			}
		}
	}
	return ans
}
