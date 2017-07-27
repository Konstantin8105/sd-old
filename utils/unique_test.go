package utils_test

import (
	"math/rand"
	"testing"

	"github.com/Konstantin8105/GoFea/solver/dof"
	"github.com/Konstantin8105/GoFea/utils"
)

func TestUniqueInt1(t *testing.T) {
	array := []int{1, 2, 3, 4, 5}
	result := []int{1, 2, 3, 4, 5}

	utils.UniqueInt(&array)
	Compare(t, array, result)
}

func TestUniqueInt2(t *testing.T) {
	array := []int{1, 1, 1, 1, 5}
	result := []int{1, 5}

	utils.UniqueInt(&array)
	Compare(t, array, result)
}

func TestUniqueInt3(t *testing.T) {
	array := []int{1, 2, 2, 2, 5}
	result := []int{1, 2, 5}

	utils.UniqueInt(&array)
	Compare(t, array, result)
}

func TestUniqueInt4(t *testing.T) {
	array := []int{1, 2, 2, 2, 2}
	result := []int{1, 2}

	utils.UniqueInt(&array)
	Compare(t, array, result)
}

func TestUniqueInt5(t *testing.T) {
	array := []int{2, 2, 2, 2, 2}
	result := []int{2}

	utils.UniqueInt(&array)
	Compare(t, array, result)
}

func TestUniqueInt6(t *testing.T) {
	array := []int{20, 200, 2, 2, 5}
	result := []int{2, 5, 20, 200}

	utils.UniqueInt(&array)
	Compare(t, array, result)
}

func TestUniqueInt7(t *testing.T) {
	array := []int{1, 200, 2, 2, 5}
	result := []int{1, 2, 5, 200}

	utils.UniqueInt(&array)
	Compare(t, array, result)
}

func TestUniqueInt8(t *testing.T) {
	array := []int{200, 200, 2, 2, 5}
	result := []int{2, 5, 200}

	utils.UniqueInt(&array)
	Compare(t, array, result)
}

func TestUniqueInt9(t *testing.T) {
	array := []int{20, 5, 2, 2, 5}
	result := []int{2, 5, 20}

	utils.UniqueInt(&array)
	Compare(t, array, result)
}

func TestUniqueInt10(t *testing.T) {
	array := []int{-1, 200, 2, 2, 5}
	result := []int{-1, 2, 5, 200}

	utils.UniqueInt(&array)
	Compare(t, array, result)
}

func Compare(t *testing.T, a, b []int) {
	if len(a) != len(b) {
		t.Errorf("Wrong")
	}
	for i := range a {
		if a[i] != b[i] {
			t.Errorf("Wrong")
		}
	}
}

// Benchmark
// Result:
// go test -bench=.  -benchtime=10s -cpu 2
// 21 may 2015:
// Benchmark1000-2   	  100000	    147824 ns/op	      32 B/op	       1 allocs/op

func Benchmark1000(b *testing.B) {
	size := 1000
	array := make([]int, size, size)
	for i := 0; i < size; i++ {
		array[i] = rand.Intn(size)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c := make([]int, size, size)
		copy(c, array)
		b.StartTimer()
		utils.UniqueInt(&c)
		b.StopTimer()
	}
	b.ReportAllocs()
}

func TestUniqueAxeNumber(t *testing.T) {
	array := []dof.AxeNumber{5, 3, 3, 5}
	result := []dof.AxeNumber{3, 5}

	utils.UniqueAxeNumber(&array)

	if len(array) != len(result) {
		t.Errorf("Wrong")
	}
	for i := range array {
		if array[i] != result[i] {
			t.Errorf("Wrong")
		}
	}
}
