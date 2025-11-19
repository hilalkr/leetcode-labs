package keepmultiplyingfoundvaluesbytwo

import "testing"

func TestFindFinalValue(t *testing.T) {
	tests := []struct {
		nums     []int
		original int
		want     int
	}{
		{nums: []int{5, 3, 6, 1, 12}, original: 3, want: 24},
		{nums: []int{2, 4, 8}, original: 3, want: 3},
		{nums: []int{2, 4, 8}, original: 2, want: 16},
	}

	for i, tt := range tests {
		got := findFinalValue(tt.nums, tt.original)
		if got != tt.want {
			t.Fatalf("test %d failed: findFinalValue(%v, %d) = %d, want %d",
				i, tt.nums, tt.original, got, tt.want)
		}
	}
}
