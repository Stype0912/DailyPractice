package accounts_merge

import (
	"sort"
)

func accountsMerge(accounts [][]string) [][]string {
	emailToIndex := map[string]int{}
	emailToName := map[string]string{}
	for _, account := range accounts {
		for _, email := range account[1:] {
			if _, has := emailToIndex[email]; !has {
				emailToIndex[email] = len(emailToIndex)
				emailToName[email] = account[0]
			}
		}
	}
	n := len(emailToIndex)
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}
	union := func(x, y int) {
		parent[find(x)] = find(y)
	}
	for _, account := range accounts {
		firstIndex := emailToIndex[account[1]]
		for _, email := range account[2:] {
			union(emailToIndex[email], firstIndex)
		}
	}
	indexToEmails := map[int][]string{}
	for email, index := range emailToIndex {
		indexToEmails[find(index)] = append(indexToEmails[find(index)], email)
	}
	ans := [][]string{}
	for _, emails := range indexToEmails {
		sort.Strings(emails)
		account := append([]string{emailToName[emails[0]]}, emails...)
		ans = append(ans, account)
	}
	return ans
}
