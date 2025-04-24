package exercises

import "testing"

func TestSumToN(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{5, 15},     // 1+2+3+4+5
		{10, 55},    // 1+2+...+10
		{0, 0},      // граничный случай
		{1, 1},      // одно число
		{100, 5050}, // большой диапазон
	}

	for _, tt := range tests {
		result := sumToN(tt.n)
		if result != tt.expected {
			t.Errorf("sumToN(%d): ожидали %d, получили %d", tt.n, tt.expected, result)
		}
	}
}
