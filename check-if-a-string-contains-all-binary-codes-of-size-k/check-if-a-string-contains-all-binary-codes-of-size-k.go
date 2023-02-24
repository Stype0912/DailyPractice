package check_if_a_string_contains_all_binary_codes_of_size_k

import (
	"fmt"
	"math"
	"strconv"
)

func hasAllCodes(s string, k int) bool {
	strMap := map[string]bool{}
	for i := 0; i < len(s)-k+1; i++ {
		strMap[s[i:i+k]] = true
	}
	length := int(math.Pow(float64(2), float64(k)))
	for i := 0; i < length; i++ {
		subStr := fmt.Sprintf("%0"+strconv.Itoa(k)+"b", i)
		if !strMap[subStr] {
			return false
		}
	}
	return true
}
