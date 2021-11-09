package main

import (
	"fmt"
	"sort"
	"strconv"
)

type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

type intList []int

func (m intList) Len() int {
	return len(m)
}

func (m intList) Less(i, j int) bool {
	return m[i] > m[j]
}

func (m intList) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func largestMultipleOfThree(digits []int) string {
	sort.Sort(intList(digits))
	sum := 0
	ans := ""
	var onePos, twoPos []int
	for i, item := range digits {
		if item%3 == 1 {
			onePos = append(onePos, i)
		}
		if item%3 == 2 {
			twoPos = append(twoPos, i)
		}
		sum += item
		ans += strconv.FormatInt(int64(item), 10)
	}
	if sum%3 == 0 {
		return dropZero(ans)
	} else if sum%3 == 1 {
		if len(onePos) == 0 && len(twoPos) < 2 {
			return ""
		} else if len(onePos) != 0 {
			return dropZero(ans[:onePos[len(onePos)-1]] + ans[onePos[len(onePos)-1]+1:])
		} else if len(twoPos) >= 2 {
			return dropZero(ans[:twoPos[len(twoPos)-2]] + ans[twoPos[len(twoPos)-2]+1:twoPos[len(twoPos)-1]] + ans[twoPos[len(twoPos)-1]+1:])
		}

	} else if sum%3 == 2 {
		if len(twoPos) == 0 && len(onePos) < 2 {
			return ""
		} else if len(twoPos) != 0 {
			return dropZero(ans[:twoPos[len(twoPos)-1]] + ans[twoPos[len(twoPos)-1]+1:])
		} else if len(onePos) >= 2 {
			return dropZero(ans[:onePos[len(onePos)-2]] + ans[onePos[len(onePos)-2]+1:onePos[len(onePos)-1]] + ans[onePos[len(onePos)-1]+1:])
		}
	}
	return ""
}

func dropZero(ans string) string {
	if ans == "" {
		return ""
	}
	for {
		if string(ans[0]) == "0" && len(ans) != 1 {
			ans = ans[1:]
			continue
		} else {
			return ans
		}
	}
}

func main() {
	fmt.Println(largestMultipleOfThree([]int{9, 8, 6, 8, 6}))
}
