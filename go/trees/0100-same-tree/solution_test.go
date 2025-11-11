package sametree

import "testing"

func node(v int, l, r *TreeNode) *TreeNode {
	return &TreeNode{Val: v, Left: l, Right: r}
}

func TestIsSameTreeIter_Table(t *testing.T) {
	tests := []struct {
		name string
		p, q *TreeNode
		want bool
	}{
		{
			name: "both_nil",
			p:    nil,
			q:    nil,
			want: true,
		},
		{
			name: "identical_small_tree",
			p:    node(1, node(2, nil, nil), node(3, nil, nil)),
			q:    node(1, node(2, nil, nil), node(3, nil, nil)),
			want: true,
		},
		{
			name: "different_value",
			p:    node(1, node(2, nil, nil), node(3, nil, nil)),
			q:    node(1, node(2, nil, nil), node(4, nil, nil)),
			want: false,
		},
		{
			name: "different_shape",
			p:    node(1, node(2, nil, nil), nil),
			q:    node(1, nil, node(2, nil, nil)),
			want: false,
		},
		{
			name: "deeper_identical",
			p: node(1,
				node(2, node(4, nil, nil), node(5, nil, nil)),
				node(3, nil, node(6, nil, nil)),
			),
			q: node(1,
				node(2, node(4, nil, nil), node(5, nil, nil)),
				node(3, nil, node(6, nil, nil)),
			),
			want: true,
		},
		{
			name: "one_nil_one_not",
			p:    node(1, nil, nil),
			q:    nil,
			want: false,
		},
	}

	for _, tc := range tests {
		got := isSameTreeIterative(tc.p, tc.q)
		if got != tc.want {
			t.Fatalf("%s: got=%v want=%v", tc.name, got, tc.want)
		}
	}
}
