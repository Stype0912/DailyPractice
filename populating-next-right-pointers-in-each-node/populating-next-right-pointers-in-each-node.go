package populating_next_right_pointers_in_each_node

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	stack := []*Node{root}
	if root == nil {
		return root
	}
	for len(stack) != 0 {
		length := len(stack)
		for i := 0; i < length; i++ {
			node := stack[0]
			stack = stack[1:]
			if i != length-1 {
				node.Next = stack[0]
			}
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
		}
	}
	return root
}
