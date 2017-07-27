package dof_test

import (
	"fmt"
	"testing"

	"github.com/Konstantin8105/GoFea/solver/dof"
)

func TestFound(t *testing.T) {
	var d dof.DoF
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	_ = d.GetDoF(1)
}

func TestRemoveIndexes(t *testing.T) {
	tests := []struct {
		array, expected []dof.AxeNumber
		removed         []int
	}{
		{
			array:    []dof.AxeNumber{0, 1, 2, 35, 44},
			removed:  []int{},
			expected: []dof.AxeNumber{0, 1, 2, 35, 44},
		},
		{
			array:    []dof.AxeNumber{0, 1, 2, 35, 44},
			removed:  []int{0},
			expected: []dof.AxeNumber{1, 2, 35, 44},
		},
		{
			array:    []dof.AxeNumber{0, 1, 2, 35, 44},
			removed:  []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			expected: []dof.AxeNumber{1, 2, 35, 44},
		},
	}
	for index, test := range tests {
		t.Run(fmt.Sprintf("Remove-%v", index), func(t *testing.T) {
			dof.RemoveIndexes(&test.array, test.removed...)
			if len(test.array) != len(test.expected) {
				t.Errorf("Wrong lenght of array.\nResult = %#v\nExpected = %#v", test.array, test.expected)
			}
			for i := range test.array {
				if test.array[i] != test.expected[i] {
					t.Errorf("Not same array. \nResult = %#v\nExpected = %#v", test.array, test.expected)
				}
			}
		})
	}
}

func TestRemoveIndexesPanic(t *testing.T) {
	tests := []struct {
		array   []dof.AxeNumber
		removed []int
	}{
		{
			array:   []dof.AxeNumber{0, 1, 2, 35, 44},
			removed: []int{66},
		},
		{
			array:   []dof.AxeNumber{0, 1, 2, 35, 44},
			removed: []int{-1},
		},
		{
			array:   []dof.AxeNumber{0, 1, 2, 35, 44},
			removed: []int{5},
		},
	}
	for index, test := range tests {
		t.Run(fmt.Sprintf("Remove-%v", index), func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("The code did not panic")
				}
			}()
			dof.RemoveIndexes(&test.array, test.removed...)
		})
	}
}

func TestUniqueAxeNumber(t *testing.T) {
	array := []dof.AxeNumber{5, 3, 3, 5}
	result := []dof.AxeNumber{3, 5}

	dof.UniqueAxeNumber(&array)

	if len(array) != len(result) {
		t.Errorf("Wrong")
	}
	for i := range array {
		if array[i] != result[i] {
			t.Errorf("Wrong")
		}
	}
}
