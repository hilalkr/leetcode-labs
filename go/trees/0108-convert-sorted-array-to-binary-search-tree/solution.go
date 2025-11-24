package convertsortedarraytobinarysearchtree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	return build(nums, 0, len(nums)-1)
}

func build(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}

	mid := (left + right) / 2
	node := &TreeNode{Val: nums[mid]}
	node.Left = build(nums, left, mid-1)
	node.Right = build(nums, mid+1, right)

	return node
}
