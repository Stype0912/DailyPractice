package US1pGT

import "testing"

func TestTrie(t *testing.T) {
	obj := Constructor()
	obj.BuildDict([]string{"hello", "leetcode"})
	t.Log(obj.Search("hhhlo"))
}
