package mergesortedarray

import "testing"

func TestMergeSortedArray(t *testing.T) {
	tests := []struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
		want  []int
	}{
		{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3, []int{1, 2, 2, 3, 5, 6}},
		{[]int{1}, 1, []int{}, 0, []int{1}},
		{[]int{0}, 0, []int{1}, 1, []int{1}},
	}
	for _, tc := range tests {
		MergeSortedArray(tc.nums1, tc.m, tc.nums2, tc.n)
		for i := range tc.want {
			if tc.nums1[i] != tc.want[i] {
				t.Fatalf("got %v, want %v", tc.nums1, tc.want)
			}
		}
	}
}
