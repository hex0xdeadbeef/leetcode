package easy

import (
	"strconv"
	"testing"
)

func Test_strStr(t *testing.T) {
	tests := []struct {
		haystack string
		needle   string
		expected int
	}{
		{
			haystack: "mississippi",
			needle:   "issip",
			expected: 4,
		},
	}

	for id, tc := range tests {
		t.Run(strconv.Itoa(id), func(t *testing.T) {
			if res := strStr(tc.haystack, tc.needle); res != tc.expected {
				t.Errorf("f(%s, %s) = %d, expected = %d", tc.haystack, tc.needle, res, tc.expected)
			}
		})
	}
}
