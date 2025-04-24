package exercises

import (
	"reflect"
	"testing"
)

func TestMergeMaps(t *testing.T) {
	m1 := map[string]int{"apple": 5, "banana": 3}
	m2 := map[string]int{"banana": 2, "orange": 4}

	expected := map[string]int{"apple": 5, "banana": 5, "orange": 4}
	result := mergeMaps(m1, m2)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("mergeMaps: ожидали %v, получили %v", expected, result)
	}
}
