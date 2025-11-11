package sametree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Recursive
func isSameTreeRecursive(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTreeRecursive(p.Left, q.Left) && isSameTreeRecursive(p.Right, q.Right)
}

//Iterative

func isSameTreeIterative(p *TreeNode, q *TreeNode) bool {
	type pair struct{ a, b *TreeNode }
	queue := []pair{{p, q}} // FIFO

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		a, b := cur.a, cur.b
		if a == nil && b == nil {
			continue
		}
		if a == nil || b == nil {
			return false
		}
		if a.Val != b.Val {
			return false
		}
		queue = append(queue, pair{a.Left, b.Left})
		queue = append(queue, pair{a.Right, b.Right})
	}
	return true
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	return isSameTreeRecursive(p, q)
}
