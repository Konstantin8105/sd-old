package utils_test

import (
	"fmt"
	"testing"

	"github.com/Konstantin8105/GoFea/utils"
)

func TestRemove(t *testing.T) {
	tests := []struct {
		a, b, expected []int
	}{
		{
			a:        []int{1, 2, 3, 4, 5},
			b:        []int{2, 3},
			expected: []int{1, 4, 5},
		},
		{
			a:        []int{1, 2, 3, 4, 5},
			b:        []int{9, 8},
			expected: []int{1, 2, 3, 4, 5},
		},
	}
	for index, test := range tests {
		t.Run(fmt.Sprintf("Remove-%v", index), func(t *testing.T) {
			result := utils.Remove(test.a, test.b)
			if len(result) != len(test.expected) {
				t.Errorf("Wrong lenght of array.\nResult = %#v\nExpected = %#v", result, test.expected)
			}
			for i := range result {
				if result[i] != test.expected[i] {
					t.Errorf("Not same array. \nResult = %#v\nExpected = %#v", result, test.expected)
				}
			}
		})
	}
}
