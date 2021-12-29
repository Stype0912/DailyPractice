package main

func minCharacters(a string, b string) int {
	MapA := make([]int, 26)
	MapB := make([]int, 26)
	for i := 0; i < len(a); i++ {
		MapA[a[i]-97]++
	}
	for i := 0; i < len(b); i++ {
		MapB[b[i]-97]++
	}
	ans := len(a)
	for i := 0; i < 26; i++ {
		way1 := Sum(MapA[i:]) + Sum(MapB[:i])
		way2 := Sum(MapB[i:]) + Sum(MapA[:i])
		way3 := len(a) - MapA[i] + len(b) - MapB[i]
		if i == 0 {
			ans = way3
		} else {
			ans = Min(ans, Min(way3, Min(way2, way1)))
		}
	}
	return ans
}

//func Min(a, b int) int {
//	if a < b {
//		return a
//	}
//	return b
//}

func Sum(list []int) int {
	ans := 0
	for _, item := range list {
		ans += item
	}
	return ans
}
