package middle

import "testing"

func Test_characterReplacementRepeat(t *testing.T) {
	tt := []struct {
		s        string
		k        int
		expected int
	}{
		{
			s:        "ABAB",
			k:        2,
			expected: 4,
		},
		{
			s:        "AABABBA",
			k:        1,
			expected: 4,
		},
		{
			s:        "ABAB",
			k:        0,
			expected: 1,
		},
		{
			s:        "ABCDE",
			k:        1,
			expected: 2,
		},
	}

	for _, tc := range tt {
		if res := characterReplacementRepeat(tc.s, tc.k); res != tc.expected {
			t.Errorf("f(%q, %d) = %d, wanted = %d", tc.s, tc.k, res, tc.expected)
		}
	}

}
