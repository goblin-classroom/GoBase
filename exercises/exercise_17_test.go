package exercises

import (
	"testing"
)

func TestSafeFunction(t *testing.T) {
	// Флаг для проверки выполнения
	called := false

	safeFunction(func() {
		called = true
		panic("что-то пошло не так")
	})

	if !called {
		t.Errorf("safeFunction: функция не была вызвана")
	}

	// Проверим, что безопасный вызов без паники тоже работает
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("safeFunction: неожиданная паника: %v", r)
		}
	}()
	safeFunction(func() {
		// ничего не делаем
	})
}
