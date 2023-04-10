package next_greater_node_in_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func nextLargerNodes(head *ListNode) []int {
	nums := []int{}
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}
	ans := make([]int, len(nums))
	stack := []int{}
	for i := len(nums) - 1; i >= 0; i-- {
		for len(stack) > 0 && stack[len(stack)-1] <= nums[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			ans[i] = 0
		} else {
			ans[i] = stack[len(stack)-1]
		}
		stack = append(stack, nums[i])
	}
	return ans
}
