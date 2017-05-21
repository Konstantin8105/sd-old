package utils

import (
	"math/rand"
	"testing"
)

func TestUniqueInt1(t *testing.T) {
	array := []int{1, 2, 3, 4, 5}
	result := []int{1, 2, 3, 4, 5}

	UniqueInt(&array)
	Compare(t, array, result)
}

func TestUniqueInt2(t *testing.T) {
	array := []int{1, 1, 1, 1, 5}
	result := []int{1, 5}

	UniqueInt(&array)
	Compare(t, array, result)
}

func TestUniqueInt3(t *testing.T) {
	array := []int{1, 2, 2, 2, 5}
	result := []int{1, 2, 5}

	UniqueInt(&array)
	Compare(t, array, result)
}

func TestUniqueInt4(t *testing.T) {
	array := []int{1, 2, 2, 2, 2}
	result := []int{1, 2}

	UniqueInt(&array)
	Compare(t, array, result)
}

func TestUniqueInt5(t *testing.T) {
	array := []int{2, 2, 2, 2, 2}
	result := []int{2}

	UniqueInt(&array)
	Compare(t, array, result)
}

func TestUniqueInt6(t *testing.T) {
	array := []int{20, 200, 2, 2, 5}
	result := []int{2, 5, 20, 200}

	UniqueInt(&array)
	Compare(t, array, result)
}

func TestUniqueInt7(t *testing.T) {
	array := []int{1, 200, 2, 2, 5}
	result := []int{1, 2, 5, 200}

	UniqueInt(&array)
	Compare(t, array, result)
}

func TestUniqueInt8(t *testing.T) {
	array := []int{200, 200, 2, 2, 5}
	result := []int{2, 5, 200}

	UniqueInt(&array)
	Compare(t, array, result)
}

func TestUniqueInt9(t *testing.T) {
	array := []int{20, 5, 2, 2, 5}
	result := []int{2, 5, 20}

	UniqueInt(&array)
	Compare(t, array, result)
}

func TestUniqueInt10(t *testing.T) {
	array := []int{-1, 200, 2, 2, 5}
	result := []int{-1, 2, 5, 200}

	UniqueInt(&array)
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
		UniqueInt(&c)
		b.StopTimer()
	}
	b.ReportAllocs()
}
