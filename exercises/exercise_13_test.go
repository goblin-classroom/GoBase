package exercises

import "testing"

func TestIncrement(t *testing.T) {
	n := 5
	increment(&n)

	if n != 6 {
		t.Errorf("increment(&n): ожидали 6, получили %d", n)
	}

	increment(&n)
	if n != 7 {
		t.Errorf("increment(&n): ожидали 7, получили %d", n)
	}
}
