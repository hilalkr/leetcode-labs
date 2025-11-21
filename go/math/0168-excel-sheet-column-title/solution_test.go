package excelsheetcolumntitle

import "testing"

func TestConvertToTitle(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  string
	}{
		{name: "A", input: 1, want: "A"},
		{name: "Z", input: 26, want: "Z"},
		{name: "AA", input: 27, want: "AA"},
		{name: "AB", input: 28, want: "AB"},
		{name: "ZY", input: 701, want: "ZY"},
		{name: "ZZ", input: 702, want: "ZZ"},
		{name: "AAA", input: 703, want: "AAA"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ConvertToTitle(tt.input)

			if got != tt.want {
				t.Errorf("ConvertToTitle(%d) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}
