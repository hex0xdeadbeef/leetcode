package easy

import "testing"

func TestFindJudge(t *testing.T) {
	tests := []struct {
		n        int
		trust    [][]int
		expected int
	}{

		{4, [][]int{{1, 3}, {1, 4}, {2, 3}}, -1},
	}

	for _, test := range tests {
		result := findJudge(test.n, test.trust)
		if result != test.expected {
			t.Errorf("findJudge(%d, %v) = %d; want %d", test.n, test.trust, result, test.expected)
		}
	}
}
