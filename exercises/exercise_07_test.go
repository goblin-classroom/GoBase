package exercises

import "testing"

func TestCalculator(t *testing.T) {
	tests := []struct {
		a, b     int
		op       string
		expected int
	}{
		{10, 5, "+", 15},
		{10, 5, "-", 5},
		{10, 5, "*", 50},
		{10, 5, "/", 2},
		{10, 0, "/", 0},   // деление на ноль
		{3, 4, "?", 0},    // неизвестная операция
		{-3, -2, "+", -5}, // отрицательные числа
	}

	for _, tt := range tests {
		result := calculator(tt.a, tt.b, tt.op)
		if result != tt.expected {
			t.Errorf("calculator(%d, %d, %q): ожидали %d, получили %d", tt.a, tt.b, tt.op, tt.expected, result)
		}
	}
}
