package binarytreeinordertraversal

import (
	"reflect"
	"testing"
)

func TestInorderTraversal_Table(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want []int
	}{
		{
			name: "example_1_[1,nil,2,3]_=>_[1,3,2]",
			//  1
			//   \
			//    2
			//   /
			//  3
			root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 3},
				},
			},
			want: []int{1, 3, 2},
		},
		{
			name: "single_node",
			root: &TreeNode{Val: 42},
			want: []int{42},
		},
		{
			name: "empty",
			root: nil,
			want: []int{},
		},
		{
			name: "bigger_tree",
			// [1,2,3,4,5,nil,8,nil,nil,6,7,9]
			//         1
			//       /   \
			//      2     3
			//     / \     \
			//    4   5     8
			//       / \   /
			//      6   7 9
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 4},
					Right: &TreeNode{
						Val:   5,
						Left:  &TreeNode{Val: 6},
						Right: &TreeNode{Val: 7},
					},
				},
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val:  8,
						Left: &TreeNode{Val: 9},
					},
				},
			},
			want: []int{4, 2, 6, 5, 7, 1, 3, 9, 8},
		},
	}

	for _, tc := range tests {
		got := inorderTraversal(tc.root)
		if !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("%s: got=%v want=%v", tc.name, got, tc.want)
		}
	}
}
