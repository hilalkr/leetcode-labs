package symmetrictree

import "testing"

func TestIsSymmetric_EmptyTree(t *testing.T) {
	if !isSymmetric(nil) {
		t.Errorf("expected true for empty tree, got false")
	}
}

func TestIsSymmetric_SymmetricTree(t *testing.T) {
	root := &TreeNode{Val: 1}

	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 2}

	root.Left.Left = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 4}

	root.Right.Left = &TreeNode{Val: 4}
	root.Right.Right = &TreeNode{Val: 3}

	if !isSymmetric(root) {
		t.Errorf("expected true for symmetric tree, got false")
	}
}

func TestIsSymmetric_NotSymmetricTree(t *testing.T) {
	root := &TreeNode{Val: 1}

	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 2}

	root.Left.Right = &TreeNode{Val: 3}
	root.Right.Right = &TreeNode{Val: 3}

	if isSymmetric(root) {
		t.Errorf("expected false for non-symmetric tree, got true")
	}
}
