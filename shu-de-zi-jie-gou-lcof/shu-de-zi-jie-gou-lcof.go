package shu_de_zi_jie_gou_lcof

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil && B == nil {
		return true
	}
	if A == nil && B != nil {
		return false
	}
	if A != nil && B == nil {
		return false
	}
	if A.Val == B.Val && isSame(A, B) {
		return true
	}
	return isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

func isSame(A, B *TreeNode) bool {
	if A == nil && B == nil {
		return true
	} else if A != nil && B != nil {
		return A.Val == B.Val && isSame(A.Left, B.Left) && isSame(A.Right, B.Right)
	} else if A != nil && B == nil {
		return true
	} else {
		return false
	}
}
