package package_problem

import "testing"

func TestPackage(t *testing.T) {
	t.Log(FindBestPackage([]int{8, 3, 4, 3}, []int{9, 3, 4, 3}, 10))
}
