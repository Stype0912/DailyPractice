package main

import (
	"sort"
	"strconv"
	"strings"
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
	var onePos []int
	twoPos := -1
	for i, item := range digits {
		if item%3 == 1 {
			onePos = append(onePos, i)
		}
		if item%3 == 2 {
			twoPos = i
		}
		sum += item
		ans += strconv.FormatInt(int64(item), 10)
	}
	if sum%3 == 0 {
		return dropZero(ans)
	} else if sum%3 == 1 {
		if len(onePos) == 0 {
			return ""
		}
		return dropZero(ans[:onePos[len(onePos)-1]] + ans[onePos[len(onePos)-1]+1:])
	} else if sum%3 == 2 {
		if twoPos == -1 && len(onePos) < 2 {
			return ""
		} else if twoPos != -1 {
			return dropZero(ans[:twoPos] + ans[twoPos+1:])
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
	slice := strings.Split(ans, "")
	for {
		if slice[0] == "0" && len(slice) != 1{
			slice = slice[1:]
			continue
		} else {
			return strings.Join(slice, "")
		}
	}
}

func main() {
	largestMultipleOfThree([]int{0, 0, 0, 0})
}
