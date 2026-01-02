package math

import "testing"

func TestMajorityElement(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Example from LeetCode - simple case",
			nums:     []int{3, 2, 3},
			expected: 3,
		},
		{
			name:     "All same elements",
			nums:     []int{2, 2, 2, 2, 2},
			expected: 2,
		},
		{
			name:     "Longer array with majority at beginning",
			nums:     []int{1, 1, 2, 3, 1, 4, 1, 5, 1},
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MajorityElement(tt.nums)
			if result != tt.expected {
				t.Errorf("MajorityElement(%v) = %d; want %d", tt.nums, result, tt.expected)
			}
		})
	}
}

func TestRepeatedNTimes(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Example from LeetCode - 2N array size 4",
			nums:     []int{1, 2, 3, 3},
			expected: 3,
		},
		{
			name:     "Larger array with repeated element",
			nums:     []int{5, 1, 5, 2, 5, 3, 5, 4},
			expected: 5,
		},
		{
			name:     "Repeated element at start",
			nums:     []int{2, 1, 2, 3, 2},
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RepeatedNTimes(tt.nums)
			if result != tt.expected {
				t.Errorf("RepeatedNTimes(%v) = %d; want %d", tt.nums, result, tt.expected)
			}
		})
	}
}
