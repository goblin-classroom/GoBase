package exercises

import (
	"reflect"
	"testing"
)

func TestAppendValue(t *testing.T) {
	s := []int{1, 2, 3}
	appendValue(&s, 4)

	expected := []int{1, 2, 3, 4}
	if !reflect.DeepEqual(s, expected) {
		t.Errorf("appendValue(&s, 4): ожидали %v, получили %v", expected, s)
	}

	appendValue(&s, 5)
	expected = []int{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(s, expected) {
		t.Errorf("appendValue(&s, 5): ожидали %v, получили %v", expected, s)
	}
}
