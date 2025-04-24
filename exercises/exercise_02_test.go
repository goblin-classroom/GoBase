package exercises

import (
	"reflect"
	"testing"
)

func TestMergeAndSort(t *testing.T) {
	tests := []struct {
		name     string
		nums1    []int
		nums2    []int
		expected []int
	}{
		{
			name:     "обычные числа",
			nums1:    []int{5, 3, 1},
			nums2:    []int{6, 4, 2},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "один пустой",
			nums1:    []int{},
			nums2:    []int{3, 1},
			expected: []int{1, 3},
		},
		{
			name:     "оба пустые",
			nums1:    []int{},
			nums2:    []int{},
			expected: []int{},
		},
		{
			name:     "отрицательные числа",
			nums1:    []int{-1, -3},
			nums2:    []int{-2, 0},
			expected: []int{-3, -2, -1, 0},
		},
		{
			name:     "повторяющиеся числа",
			nums1:    []int{1, 2, 2},
			nums2:    []int{2, 3},
			expected: []int{1, 2, 2, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := mergeAndSort(tt.nums1, tt.nums2)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("ожидали %v, получили %v", tt.expected, result)
			}
		})
	}
}
