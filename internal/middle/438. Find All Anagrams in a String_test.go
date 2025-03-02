package middle

import (
	"slices"
	"testing"
)

func Test_findAnagrams(t *testing.T) {
	tt := []struct {
		s, p     string
		expected []int
	}{
		{
			s:        "cbaebabacd",
			p:        "abc",
			expected: []int{0, 6},
		},

		{
			s:        "abab",
			p:        "ab",
			expected: []int{0, 1, 2},
		},
		{
			s:        "aa",
			p:        "bb",
			expected: nil,
		},
	}

	for _, tc := range tt {
		if res := findAnagramsDiff(tc.s, tc.p); !slices.Equal(tc.expected, res) {
			t.Errorf("f(%q %q) = %v, wanted: %v", tc.s, tc.p, res, tc.expected)
		}
	}
}
