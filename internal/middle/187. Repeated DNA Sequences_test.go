package middle

import (
	"slices"
	"testing"
)

func Test_findRepeatedDnaSequences(t *testing.T) {

	tt := []struct {
		s        string
		expected []string
	}{
		{
			s:        "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT",
			expected: []string{"AAAAACCCCC", "CCCCCAAAAA"},
		},
		{
			s:        "AAAAAAAAAAAAA",
			expected: []string{"AAAAAAAAAA"},
		},
		{
			s:        "AAAAAAAAAAA",
			expected: []string{"AAAAAAAAAA"},
		},
	}

	for _, tc := range tt {

		res := findRepeatedDnaSequences(tc.s)
		slices.Sort(tc.expected)
		slices.Sort(res)

		if !slices.Equal(tc.expected, res) {
			t.Errorf("f(%q) = %v, wanted = %v", tc.s, res, tc.expected)
		}
	}
}
