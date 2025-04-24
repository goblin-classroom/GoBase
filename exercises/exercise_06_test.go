package exercises

import "testing"

func TestCharCase(t *testing.T) {
	tests := []struct {
		input    rune
		expected string
	}{
		{'A', "Латинская заглавная"},
		{'z', "Латинская строчная"},
		{'Б', "Кириллическая заглавная"},
		{'ю', "Кириллическая строчная"},
		{'9', "Другое"},
		{'😀', "Другое"},
		{'ё', "Кириллическая строчная"}, // важно: в Go 'ё' — отдельная руна
		{'Ё', "Кириллическая заглавная"},
	}

	for _, tt := range tests {
		result := charCase(tt.input)
		if result != tt.expected {
			t.Errorf("charCase(%q): ожидали %q, получили %q", tt.input, tt.expected, result)
		}
	}
}
