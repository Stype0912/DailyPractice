package best_team_with_no_conflicts

import (
	"sort"
)

type People [][]int

func (p People) Len() int {
	return len(p[0])
}

func (p People) Less(i, j int) bool {
	if p[1][i] < p[1][j] {
		return true
	} else if p[1][i] == p[1][j] {
		if p[0][i] < p[0][j] {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func (p People) Swap(i, j int) {
	p[0][i], p[0][j] = p[0][j], p[0][i]
	p[1][i], p[1][j] = p[1][j], p[1][i]
}

func bestTeamScore(scores []int, ages []int) int {
	people := [][]int{}
	people = append(people, scores)
	people = append(people, ages)
	sort.Sort(People(people))
	dp := make([]int, len(ages))
	dp[0] = people[0][0]
	for i := 1; i < len(ages); i++ {
		for j := 0; j < i; j++ {
			if !(people[0][i] < people[0][j] && people[1][i] > people[1][j]) {
				dp[i] = max(dp[i], dp[j])
			}
		}
		dp[i] += people[0][i]
	}
	ans := 0
	for i := 0; i < len(ages); i++ {
		ans = max(ans, dp[i])
	}
	return ans
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
