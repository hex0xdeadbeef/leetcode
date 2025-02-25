package middle

import "testing"

func Test_lengthOfLongestSubstringRepeat(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			input:    "abcabcbb",
			expected: 3,
		},
		{
			input:    "bbbbb",
			expected: 1,
		},
		{
			input:    "pwwkew",
			expected: 3,
		},
		{
			input:    " ",
			expected: 1,
		},
		{
			input:    "au",
			expected: 2,
		},
		{
			input:    "aab",
			expected: 2,
		},
		{
			input:    "dvdf",
			expected: 3,
		},
		{
			input:    "abba",
			expected: 2,
		},
		{
			input:    "wobgrovw",
			expected: 6,
		},
	}

	for _, tc := range tests {
		if res := lengthOfLongestSubstringRepeat(tc.input); res != tc.expected {
			t.Errorf(" f(%q) = %d, Expected: %d", tc.input, res, tc.expected)
		}
	}
}
