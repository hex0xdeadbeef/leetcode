package easy

import "testing"

func TestIsIsomorphic(t *testing.T) {
	tests := []struct {
		s, t     string
		expected bool
	}{
		{
			s:        "aab",
			t:        "aaa",
			expected: false,
		},
	}

	for _, test := range tests {
		result := isIsomorphic(test.s, test.t)
		if result != test.expected {
			t.Errorf("isIsomorphic(%q, %q) = %v; want %v", test.s, test.t, result, test.expected)
		}
	}
}
