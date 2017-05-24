package linearAlgebra

import "fmt"

// Vector - vector data structure
type Vector struct {
	size     int
	capacity int
	values   []float64
}

// NewVector - constructor for vector
func NewVector(size int) (v Vector) {
	v.SetSize(size)
	return
}

// SetSize - change size of vector
func (v *Vector) SetSize(size int) {
	v.size = size
	if v.capacity < size {
		v.values = make([]float64, size, size)
		v.capacity = size
		return
	}
	// initialize the vector
	for i := 0; i < size; i++ {
		v.values[i] = 0.0
	}
}

// Set - change value of vector
func (v *Vector) Set(i int, value float64) {
	if i < 0 || i >= v.size {
		panic(fmt.Errorf("Cannot index is outside of vector [%v]\nVector - %#v", i, v))
	}
	v.values[i] = value
}

// Get - return value of vector
func (v *Vector) Get(i int) float64 {
	if i < 0 || i >= v.size {
		panic(fmt.Errorf("Cannot index is outside of vector [%v]\nVector - %#v", i, v))
	}
	return v.values[i]
}

func (v Vector) String() (s string) {
	for j := 0; j < v.size; j++ {
		s += fmt.Sprintf("[")
		s += fmt.Sprintf("%+.2E", v.values[j])
		s += fmt.Sprintf("]\n")
	}
	return s
}
