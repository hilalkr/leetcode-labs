package sum

import (
	"sort"
	"testing"
)

func TestThreeSum(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "example case",
			nums: []int{-1, 0, 1, 2, -1, -4},
			want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name: "all zeros collapse to single triple",
			nums: []int{0, 0, 0, 0},
			want: [][]int{{0, 0, 0}},
		},
		{
			name: "multiple duplicates still unique",
			nums: []int{-2, 0, 1, 1, 2},
			want: [][]int{{-2, 0, 2}, {-2, 1, 1}},
		},
		{
			name: "no result when sum never zero",
			nums: []int{0, 1, 1},
			want: nil,
		},
		{
			name: "input shorter than three numbers",
			nums: []int{1, 2},
			want: nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			nums := append([]int(nil), tt.nums...)
			got := threeSum(nums)
			if !equalTriples(got, tt.want) {
				t.Fatalf("threeSum(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}

func equalTriples(got, want [][]int) bool {
	normGot := normalizeTriples(got)
	normWant := normalizeTriples(want)
	if len(normGot) != len(normWant) {
		return false
	}
	for i := range normGot {
		if len(normGot[i]) != len(normWant[i]) {
			return false
		}
		for j := range normGot[i] {
			if normGot[i][j] != normWant[i][j] {
				return false
			}
		}
	}
	return true
}

func normalizeTriples(src [][]int) [][]int {
	dst := make([][]int, len(src))
	for i, triple := range src {
		cp := append([]int(nil), triple...)
		sort.Ints(cp)
		dst[i] = cp
	}
	sort.Slice(dst, func(i, j int) bool {
		ai, aj := dst[i], dst[j]
		for k := 0; k < len(ai) && k < len(aj); k++ {
			if ai[k] != aj[k] {
				return ai[k] < aj[k]
			}
		}
		return len(ai) < len(aj)
	})
	return dst
}
