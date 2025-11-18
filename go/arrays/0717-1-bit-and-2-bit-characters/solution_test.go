package bitand2bitcharacters

import "testing"

func TestIsOneBitCharacter(t *testing.T) {
	tests := []struct {
		bits []int
		want bool
	}{
		{[]int{1, 0, 0}, true},
		{[]int{1, 1, 1, 0}, false},
		{[]int{0}, true},
		{[]int{1, 0}, false},
		{[]int{0, 0}, true},
	}

	for _, tt := range tests {
		got := isOneBitCharacter(tt.bits)
		if got != tt.want {
			t.Errorf("isOneBitCharacter(%v) = %v, want %v", tt.bits, got, tt.want)
		}
	}
}
