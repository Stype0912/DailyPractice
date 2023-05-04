package word_search

import "testing"

func TestWord(t *testing.T) {
	t.Log(exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'E', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCESEEEFS"))
}
