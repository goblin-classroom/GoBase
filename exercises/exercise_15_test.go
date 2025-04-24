package exercises

import (
	"reflect"
	"testing"
)

func TestRemoveAtIndex(t *testing.T) {
	tests := []struct {
		arr      []int
		index    int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, []int{1, 2, 4, 5}},
		{[]int{10, 20, 30}, 5, []int{10, 20, 30}}, // индекс вне границ
		{[]int{1}, 0, []int{}},
		{[]int{}, 0, []int{}},
		{[]int{7, 8, 9}, -1, []int{7, 8, 9}}, // отрицательный индекс
	}

	for _, tt := range tests {
		result := removeAtIndex(tt.arr, tt.index)
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("removeAtIndex(%v, %d): ожидали %v, получили %v", tt.arr, tt.index, tt.expected, result)
		}
	}
}
