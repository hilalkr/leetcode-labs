package binarytreeinordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Recursive
func InorderTraversalRecursive(root *TreeNode) []int {
	res := make([]int, 0)

	var walk func(*TreeNode)
	walk = func(node *TreeNode) {
		if node == nil {
			return
		}
		walk(node.Left)
		res = append(res, node.Val)
		walk(node.Right)
	}
	walk(root)
	return res
}

//Iterate(with stack)

func InorderTraversalIterative(root *TreeNode) []int {
	res := make([]int, 0)
	stack := []*TreeNode{}
	curr := root

	for curr != nil || len(stack) > 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}

		n := len(stack) - 1
		curr = stack[n]
		stack = stack[:n]

		res = append(res, curr.Val)

		curr = curr.Right
	}
	return res
}

func inorderTraversal(root *TreeNode) []int {
	return InorderTraversalIterative(root)
}
