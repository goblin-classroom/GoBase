package exercises

import "testing"

func TestMaxOfThree(t *testing.T) {
	tests := []struct {
		a, b, c  int
		expected int
	}{
		{1, 2, 3, 3},
		{3, 2, 1, 3},
		{2, 3, 1, 3},
		{5, 5, 5, 5},
		{10, 10, 5, 10},
		{-1, -2, -3, -1},
	}

	for _, tt := range tests {
		result := maxOfThree(tt.a, tt.b, tt.c)
		if result != tt.expected {
			t.Errorf("maxOfThree(%d, %d, %d): ожидали %d, получили %d", tt.a, tt.b, tt.c, tt.expected, result)
		}
	}
}
