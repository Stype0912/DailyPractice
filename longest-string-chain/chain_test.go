package longest_string_chain

import "testing"

func TestChain(t *testing.T) {
	t.Log(longestStrChain([]string{"a", "b", "ba", "bca", "bda", "bdca"}))
}
