package exercises

import "testing"

func TestSequenceGenerator(t *testing.T) {
	next := sequenceGenerator()

	values := []int{1, 2, 3, 4, 5}
	for _, expected := range values {
		result := next()
		if result != expected {
			t.Errorf("sequenceGenerator(): ожидали %d, получили %d", expected, result)
		}
	}
}
