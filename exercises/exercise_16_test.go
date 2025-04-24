package exercises

import "testing"

func TestSafeDivide(t *testing.T) {
	tests := []struct {
		a, b     int
		expected int
	}{
		{10, 2, 5},
		{9, 3, 3},
		{7, 0, 0}, // деление на ноль — ожидание 0
	}

	for _, tt := range tests {
		result := safeDivide(tt.a, tt.b)
		if result != tt.expected {
			t.Errorf("safeDivide(%d, %d): ожидали %d, получили %d", tt.a, tt.b, tt.expected, result)
		}
	}
}
