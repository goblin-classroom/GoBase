package exercises

import "testing"

func TestReverseNumber(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{123, 321},
		{500, 5},
		{-987, -789},
		{0, 0},
		{-100, -1},
		{1, 1},
	}

	for _, tt := range tests {
		result := reverseNumber(tt.input)
		if result != tt.expected {
			t.Errorf("reverseNumber(%d): ожидали %d, получили %d", tt.input, tt.expected, result)
		}
	}
}
