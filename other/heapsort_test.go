// Problem:
// Implement heapsort.
//
// Mechanic:
// Similar to selection sort, repeatedly choose the largest item and move it to
// the end of the array using a heap.
//
// Cost:
// O(nlogn) time and O(1) space.

package other

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestHeapsort(t *testing.T) {
	tests := []struct {
		in       []int
		expected []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 2}, []int{1, 2}},
		{[]int{2, 1}, []int{1, 2}},
		{[]int{2, 1, 3}, []int{1, 2, 3}},
		{[]int{1, 1, 1}, []int{1, 1, 1}},
		{[]int{2, 1, 2}, []int{1, 2, 2}},
		{[]int{1, 2, 4, 3, 6, 5}, []int{1, 2, 3, 4, 5, 6}},
		{[]int{6, 2, 4, 3, 1, 5}, []int{1, 2, 3, 4, 5, 6}},
		{[]int{6, 5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5, 6}},
	}

	for _, tt := range tests {
		heapsort(tt.in)
		common.Equal(t, tt.expected, tt.in)
	}
}

func heapsort(in []int) {
	minIndex := 0
	for i := 0; i < len(in)-1; i++ {
		minIndex = i
		// find the minimum in the rest of the array.
		for j := i + 1; j < len(in); j++ {
			if in[j] < in[minIndex] {
				minIndex = j
			}
		}

		// swap the minimum value with the first value.
		tmp := in[i]
		in[i] = in[minIndex]
		in[minIndex] = tmp
	}
}