package middle

import (
	"reflect"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		nums     []int
		k        int
		expected [][]int
	}{
		{
			nums:     []int{1, 1, 1, 2, 2, 3},
			k:        2,
			expected: [][]int{{1, 2}, {2, 1}},
		},
	}

	for _, test := range tests {
		result := topKFrequent(test.nums, test.k)

		found := false
		for _, exp := range test.expected {
			if reflect.DeepEqual(result, exp) {
				found = true
				break
			}
		}

		if !found {
			t.Errorf("topKFrequent(%v, %d) = %v; want one of %v", test.nums, test.k, result, test.expected)
		}
	}
}
