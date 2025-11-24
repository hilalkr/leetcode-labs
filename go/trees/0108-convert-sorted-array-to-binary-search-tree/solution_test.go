package convertsortedarraytobinarysearchtree

import "testing"

func inorder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	inorder(root.Left, res)
	*res = append(*res, root.Val)
	inorder(root.Right, res)
}

func TestSortedArrayToBST_InorderEqualsInput(t *testing.T) {
	nums := []int{-10, -3, 0, 5, 9}

	root := sortedArrayToBST(nums)

	var got []int
	inorder(root, &got)

	if len(got) != len(nums) {
		t.Fatalf("length mismatch: got %d, want %d", len(got), len(nums))
	}

	for i, v := range nums {
		if got[i] != v {
			t.Fatalf("inorder[%d] = %d, want %d", i, got[i], v)
		}
	}
}
