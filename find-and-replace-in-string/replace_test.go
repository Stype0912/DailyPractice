package find_and_replace_in_string

import "testing"

func TestReplace(t *testing.T) {
	t.Log(findReplaceString("abcde", []int{2, 2}, []string{"cdef", "bc"}, []string{"f", "fe"}))
}
