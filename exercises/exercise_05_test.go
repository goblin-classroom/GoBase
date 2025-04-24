package exercises

import "testing"

func TestReverse(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello", "olleh"},
		{"Привет", "тевирП"},
		{"", ""},
		{"a", "a"},
		{"Go 💙", "💙 oG"},
		{"12345", "54321"},
	}

	for _, tt := range tests {
		result := reverse(tt.input)
		if result != tt.expected {
			t.Errorf("reverse(%q): ожидали %q, получили %q", tt.input, tt.expected, result)
		}
	}
}
