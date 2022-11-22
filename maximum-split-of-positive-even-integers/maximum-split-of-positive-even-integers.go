package maximum_split_of_positive_even_integers

import "math"

func maximumEvenSplit(finalSum int64) []int64 {
	var res []int64
	if finalSum%2 == 1 {
		return res
	}
	n := int64(math.Ceil(math.Sqrt(float64(4*finalSum+1)) - 1))
	if n%2 == 1 {
		n += 1
	}
	var sum int64 = 0
	for i := int64(2); i <= n; i += 2 {
		sum += i
		res = append(res, i)
	}
	toBeDeleted := sum - finalSum
	if toBeDeleted == 0 {
		return res
	} else {
		return append(res[:(toBeDeleted/2)-1], res[(toBeDeleted)/2:]...)
	}
	return res
}
