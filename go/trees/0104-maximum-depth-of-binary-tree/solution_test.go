package maximumdepthofbinarytree

import "testing"

func TestMaxDepthExample(t *testing.T) {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}

	got := maxDepth(root)
	want := 3

	if got != want {
		t.Errorf("maxDepth(example) = %d, want %d", got, want)
	}
}

func TestMaxDepthEmpty(t *testing.T) {
	got := maxDepth(nil)
	want := 0

	if got != want {
		t.Errorf("maxDepth(nil) = %d, want %d", got, want)
	}
}

func TestMaxDepthSingleNode(t *testing.T) {
	root := &TreeNode{Val: 1}

	got := maxDepth(root)
	want := 1

	if got != want {
		t.Errorf("maxDepth(single node) = %d, want %d", got, want)
	}
}
