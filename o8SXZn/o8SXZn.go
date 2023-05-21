package o8SXZn

import "math"

func storeWater(bucket []int, vat []int) int {
	n := len(bucket)
	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	min := func(i, j int) int {
		if i < j {
			return i
		}
		return j
	}
	maxK := 0
	for i := 0; i < n; i++ {
		maxK = max(maxK, vat[i])
	}
	if maxK == 0 {
		return 0
	}
	res := math.MaxInt
	for k := 1; k <= maxK && k < res; k++ {
		t := 0
		for i := 0; i < n; i++ {
			t += max(0, (vat[i]+k-1)/k-bucket[i])
		}
		res = min(res, t+k)
	}
	return res
}
