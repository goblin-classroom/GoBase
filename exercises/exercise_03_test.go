package exercises

import "testing"

func TestIsLeapYear(t *testing.T) {
	tests := []struct {
		year     int
		expected bool
	}{
		{2020, true},  // делится на 4, не делится на 100
		{1900, false}, // делится на 100, но не на 400
		{2000, true},  // делится на 400
		{2021, false}, // не делится на 4
		{2400, true},  // делится на 400
		{2100, false}, // делится на 100, но не на 400
	}

	for _, tt := range tests {
		result := isLeapYear(tt.year)
		if result != tt.expected {
			t.Errorf("isLeapYear(%d): ожидали %v, получили %v", tt.year, tt.expected, result)
		}
	}
}
