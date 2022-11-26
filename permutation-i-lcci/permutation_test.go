package permutation_i_lcci

import "testing"

func TestPermutation(t *testing.T) {
	t.Log(permutation("ab"))
	a := []int64{0, 1}
	b := [][]int64{{2, 3}, {4, 5}}
	b = append(b, a)
	t.Log(b)
	a[0] = 6
	t.Log(b)
	c := int64(0)
	d := []int64{1, 2}
	d = append(d, c)
	t.Log(d)
	c = 5
	t.Log(d)
}
