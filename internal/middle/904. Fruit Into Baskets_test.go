package middle

import (
	"testing"
)

func Test_totalFruit(t *testing.T) {
	tt := []struct {
		input    []int
		expected int
	}{
		{
			input:    []int{1, 2, 1},
			expected: 3,
		},
		{
			input:    []int{0, 1, 2, 2},
			expected: 3,
		},
		{
			input:    []int{1, 2, 3, 2, 2},
			expected: 4,
		},
		{
			input:    []int{0, 1, 6, 6, 4, 4, 6},
			expected: 5,
		},
		{
			input:    []int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4},
			expected: 5,
		},
		{
			input:    []int{5, 0, 0, 7, 0, 7, 2, 7},
			expected: 5,
		},
		{
			input:    []int{3, 3, 3, 1, 4},
			expected: 4,
		},
		{
			input:    []int{1, 0, 1, 4, 1, 4, 1, 2, 3},
			expected: 5,
		},
	}

	for _, tc := range tt {
		t.Run("", func(t *testing.T) {
			if res := totalFruit(tc.input); res != tc.expected {
				t.Errorf("f(%v) = %d, wanted = %d", tc.input, res, tc.expected)
			}
		})
	}

}
