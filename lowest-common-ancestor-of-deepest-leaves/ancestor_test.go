package lowest_common_ancestor_of_deepest_leaves

import "testing"

func TestAncestor(t *testing.T) {
	//t.Log(lcaDeepestLeaves(&TreeNode{
	//	Val: 1,
	//	Left: &TreeNode{
	//		Val:  2,
	//		Left: nil,
	//		Right: &TreeNode{
	//			Val:  4,
	//			Left: &TreeNode{Val: 15},
	//			Right: &TreeNode{
	//				Val:  5,
	//				Left: nil,
	//				Right: &TreeNode{
	//					Val: 7,
	//					Left: &TreeNode{
	//						Val:   8,
	//						Left:  nil,
	//						Right: &TreeNode{Val: 9},
	//					},
	//					Right: &TreeNode{
	//						Val:   12,
	//						Left:  &TreeNode{Val: 13},
	//						Right: &TreeNode{Val: 14},
	//					},
	//				},
	//			},
	//		},
	//	},
	//	Right: &TreeNode{
	//		Val: 3,
	//		Left: &TreeNode{
	//			Val: 6,
	//			Left: &TreeNode{
	//				Val: 10,
	//				Left: &TreeNode{
	//					Val:   11,
	//					Left:  nil,
	//					Right: nil,
	//				},
	//				Right: nil,
	//			},
	//			Right: nil,
	//		},
	//		Right: nil,
	//	},
	//}))
	t.Log(lcaDeepestLeaves(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: &TreeNode{Val: 8},
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:  4,
				Left: nil,
				Right: &TreeNode{
					Val:  5,
					Left: nil,
					Right: &TreeNode{
						Val: 7,
						Left: &TreeNode{
							Val:   9,
							Left:  &TreeNode{Val: 10},
							Right: &TreeNode{Val: 12},
						},
						Right: nil,
					},
				},
			},
			Right: &TreeNode{
				Val:   6,
				Left:  &TreeNode{Val: 11},
				Right: nil,
			},
		},
	}))
}
