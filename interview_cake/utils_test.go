package main

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestMimax(t *testing.T) {
	tests := []struct {
		in       []int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{1, 5}},
		{[]int{-1, -2, -3, -4, -5}, []int{-5, -1}},
		{[]int{-2, -1, 0, 1, 2}, []int{-2, 2}},
		{[]int{-10, -10, 1, 3, 2}, []int{-10, 3}},
		{[]int{1, 10, -5, 1, -100}, []int{-100, 10}},
	}

	for _, tt := range tests {
		min, max := mimax(tt.in...)
		result := []int{min, max}
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("should be %v instead %v", tt.expected, result)
		}
	}
}

func TestRandom(t *testing.T) {
	tests := []struct {
		in1      int
		in2      int
		expected int
	}{
		{1, 10, 5},
		{-10, 10, 8},
		{0, 10, 5},
		{4, 4, 4},
	}

	for _, tt := range tests {
		result := random(tt.in1, tt.in2)
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("should be %v instead %v", tt.expected, result)
		}
	}
}

// mimax returns min and max from a list of numbers.
func mimax(nums ...int) (int, int) {
	min, max := nums[0], nums[0]

	for _, num := range nums {
		if min > num {
			min = num
		}

		if max < num {
			max = num
		}
	}

	return min, max
}

// random rertuns a random number over a range
func random(min, max int) int {
	if min == max {
		return min
	}

	return rand.Intn(max-min) + min
}