package exercises

import "testing"

func TestPowerFunc(t *testing.T) {
	square := powerFunc(2)
	cube := powerFunc(3)

	tests := []struct {
		f        func(int) int
		input    int
		expected int
	}{
		{square, 4, 16},
		{square, 5, 25},
		{cube, 2, 8},
		{cube, 3, 27},
		{powerFunc(0), 100, 1}, // любое число в степени 0 = 1
	}

	for _, tt := range tests {
		result := tt.f(tt.input)
		if result != tt.expected {
			t.Errorf("powerFunc(...)(%d): ожидали %d, получили %d", tt.input, tt.expected, result)
		}
	}
}
