package middle

import "testing"

func Test_minSubArrayLen(t *testing.T) {
	tt := []struct {
		nums     []int
		target   int
		expected int
	}{
		{
			nums:     []int{2, 3, 1, 2, 4, 3},
			target:   7,
			expected: 2,
		},

		{
			nums:     []int{1, 4, 4},
			target:   4,
			expected: 1,
		},
		{
			nums:     []int{1, 1, 1, 1, 1, 1, 1, 1},
			target:   11,
			expected: 0,
		},
	}

	for _, tc := range tt {
		if res := minSubArrayLen(tc.target, tc.nums); res != tc.expected {
			t.Errorf("f(%d %v = %d, wanted = %d)", tc.target, tc.nums, res, tc.expected)
		}
	}

}
