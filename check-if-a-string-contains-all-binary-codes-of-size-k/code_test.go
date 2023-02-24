package check_if_a_string_contains_all_binary_codes_of_size_k

import "testing"

func TestCode(t *testing.T) {
	t.Log(hasAllCodes("00110110", 2))
}
