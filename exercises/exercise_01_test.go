package exercises

import (
	"testing"
)

func TestSecondMax(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
		wantErr  bool
	}{
		{"обычный случай", []int{10, 20, 5, 8}, 10, false},
		{"отрицательные числа", []int{-1, -5, -3}, -3, false},
		{"повторы", []int{5, 5, 5, 3}, 5, false},
		{"меньше двух", []int{42}, 0, true},
		{"два числа", []int{7, 1}, 1, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := secondMax(tt.input)

			if (err != nil) != tt.wantErr {
				t.Errorf("ожидали ошибку: %v, получили: %v", tt.wantErr, err)
				return
			}
			if !tt.wantErr && result != tt.expected {
				t.Errorf("ожидали %d, получили %d", tt.expected, result)
			}
		})
	}
}
