package find_the_array_concatenation_value

import "strconv"

func findTheArrayConcVal(nums []int) int64 {
	n := len(nums)
	i := 0
	j := n - 1
	ans := int64(0)
	conc := func(i, j int) int64 {
		tmp := strconv.FormatInt(int64(i), 10) + strconv.FormatInt(int64(j), 10)
		res, _ := strconv.ParseInt(tmp, 10, 64)
		return res
	}
	for i <= j {
		if i == j {
			ans += int64(nums[i])
			break
		}
		ans += conc(nums[i], nums[j])
		i++
		j--
	}
	return ans
}
