package middle

import "testing"

func TestMaxSatisfied(t *testing.T) {
	tests := []struct {
		customers []int
		grumpy    []int
		minutes   int
		expected  int
	}{
		{
			customers: []int{1, 0, 1, 2, 1, 1, 7, 5},
			grumpy:    []int{0, 1, 0, 1, 0, 1, 0, 1},
			minutes:   3,
			expected:  16,
		},

		{
			customers: []int{4, 10, 10},
			grumpy:    []int{1, 1, 0},
			minutes:   2,
			expected:  24,
		},
	}

	for _, tc := range tests {
		if res := MaxSatisfied(tc.customers, tc.grumpy, tc.minutes); res != tc.expected {
			t.Errorf("f(%v, %v %d) = %d, wanted = %d", tc.customers, tc.grumpy, tc.minutes, res, tc.expected)
		}
	}
}
