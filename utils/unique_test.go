package utils_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/Konstantin8105/GoFea/utils"
)

func TestUniqueInt1(t *testing.T) {
	tc := []struct {
		input, expected []int
	}{
		{
			[]int{1, 2, 3, 4, 5},
			[]int{1, 2, 3, 4, 5},
		},
		{
			[]int{1, 1, 1, 1, 5},
			[]int{1, 5},
		},
		{
			[]int{1, 2, 2, 2, 5},
			[]int{1, 2, 5},
		},
		{
			[]int{1, 2, 2, 2, 2},
			[]int{1, 2},
		},
		{
			[]int{2, 2, 2, 2, 2},
			[]int{2},
		},
		{
			[]int{20, 200, 2, 2, 5},
			[]int{2, 5, 20, 200},
		},
		{
			[]int{1, 200, 2, 2, 5},
			[]int{1, 2, 5, 200},
		},
		{
			[]int{200, 200, 2, 2, 5},
			[]int{2, 5, 200},
		},
		{
			[]int{20, 5, 2, 2, 5},
			[]int{2, 5, 20},
		},
		{
			[]int{-1, 200, 2, 2, 5},
			[]int{-1, 2, 5, 200},
		},
	}
	for index, test := range tc {
		t.Run(fmt.Sprintf("Case-%v", index), func(t *testing.T) {
			utils.UniqueInt(&test.input)
			if len(test.input) != len(test.expected) {
				t.Errorf("Wrong")
			}
			for i := range test.input {
				if test.input[i] != test.expected[i] {
					t.Errorf("Wrong")
				}
			}
		})
	}
}

// Benchmark
// Result:
// go test -bench=.  -benchtime=10s -cpu 2
// 21 may 2015:
// Benchmark1000-2   	  100000	    147824 ns/op	      32 B/op	       1 allocs/op

func Benchmark1000(b *testing.B) {
	size := 1000
	array := make([]int, size)
	for i := 0; i < size; i++ {
		array[i] = rand.Intn(size)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c := make([]int, size)
		copy(c, array)
		b.StartTimer()
		utils.UniqueInt(&c)
		b.StopTimer()
	}
	b.ReportAllocs()
}
