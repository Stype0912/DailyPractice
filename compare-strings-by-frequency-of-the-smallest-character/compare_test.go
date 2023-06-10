package compare_strings_by_frequency_of_the_smallest_character

import "testing"

func TestCompare(t *testing.T) {
	t.Log(numSmallerByFrequency([]string{"bbb", "cc"}, []string{"a", "aa", "aaa", "aaaa"}))
}
