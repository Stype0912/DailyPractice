package adding_two_negabinary_numbers

func addNegabinary(arr1 []int, arr2 []int) (ans []int) {
	i := len(arr1) - 1
	j := len(arr2) - 1
	carry := 0
	x := 0
	for i >= 0 || j >= 0 || carry != 0 {
		x = carry
		if i >= 0 {
			x += arr1[i]
		}
		if j >= 0 {
			x += arr2[i]
		}
		if x >= 2 {
			ans = append(ans, x-2)
			carry = -1
		} else if x >= 0 {
			ans = append(ans, x)
			carry = 0
		} else {
			ans = append(ans, 1)
			carry = 1
		}
		i--
		j--
	}
	for len(ans) > 1 && ans[len(ans)-1] == 0 {
		ans = ans[:len(ans)-1]
	}
	for k := 0; k < len(ans)/2; k++ {
		ans[k], ans[len(ans)-k] = ans[len(ans)-k], ans[k]
	}
	return ans
}
