package exercises

import (
	"reflect"
	"testing"
)

func TestWordCount(t *testing.T) {
	tests := []struct {
		input    string
		expected map[string]int
	}{
		{
			"Hello, hello world! Hello world.",
			map[string]int{"hello": 3, "world": 2},
		},
		{
			"Go go GO!",
			map[string]int{"go": 3},
		},
		{
			"Привет, мир! Привет мир.",
			map[string]int{"привет": 2, "мир": 2},
		},
		{
			"", // пустая строка
			map[string]int{},
		},
	}

	for _, tt := range tests {
		result := wordCount(tt.input)
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("wordCount(%q): ожидали %v, получили %v", tt.input, tt.expected, result)
		}
	}
}
