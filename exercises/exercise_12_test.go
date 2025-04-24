package exercises

import "testing"

func TestSwap(t *testing.T) {
	x, y := 10, 20
	swap(&x, &y)

	if x != 20 || y != 10 {
		t.Errorf("swap(&x, &y): ожидали x=20, y=10, получили x=%d, y=%d", x, y)
	}

	// Тест на одинаковые значения
	a, b := 5, 5
	swap(&a, &b)
	if a != 5 || b != 5 {
		t.Errorf("swap(&a, &b): ожидали a=5, b=5, получили a=%d, b=%d", a, b)
	}
}
