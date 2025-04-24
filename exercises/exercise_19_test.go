package exercises

import (
	"testing"
)

func TestInvertMap(t *testing.T) {
	input := map[string]int{"a": 1, "b": 2, "c": 1}
	result := invertMap(input)

	if len(result) != 2 {
		t.Errorf("ожидали 2 ключа, получили %d", len(result))
	}

	if result[2] != "b" {
		t.Errorf(`ожидали result[2] == "b", получили %q`, result[2])
	}

	// result[1] может быть "a" или "c" — это допустимо
	if result[1] != "a" && result[1] != "c" {
		t.Errorf(`result[1] должен быть "a" или "c", получили %q`, result[1])
	}
}
