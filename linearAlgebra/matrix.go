package linearAlgebra

import "fmt"

// Matrix - matrix data structure
type Matrix struct {
	size     int
	capacity int
	values   [][]float64
}

// NewSquareMatrix - constructor for square matrix
func NewSquareMatrix(size int) (m Matrix) {
	m = *new(Matrix)
	m.SetSize(size)
	return
}

// SetSize - change the matrix size
func (m *Matrix) SetSize(s int) {
	if m.capacity < s {
		m.values = make([][]float64, s, s)
		m.size = s
		m.capacity = s
		for i := 0; i < s; i++ {
			m.values[i] = make([]float64, s)
		}
		return
	}
	m.size = s
}

// Set - change value of matrix
func (m *Matrix) Set(i int, j int, value float64) {
	if i < 0 || i >= m.size || j < 0 || j >= m.size {
		panic(fmt.Errorf("Cannot index is outside the matrix - [%v,%v]\nMatrix - %#v", i, j, m))
	}
	m.values[i][j] = value
}

// Get - return value in matrix
func (m *Matrix) Get(i int, j int) float64 {
	if i < 0 || i >= m.size || j < 0 || j >= m.size {
		panic(fmt.Errorf("Cannot index is outside the matrix - [%v,%v]\nMatrix - %#v", i, j, m))
	}
	return m.values[i][j]
}

func (m Matrix) String() (s string) {
	for i := 0; i < m.size; i++ {
		s += fmt.Sprintf("[")
		for j := 0; j < m.size; j++ {
			s += fmt.Sprintf("%+.2E,", m.values[i][j])
		}
		s += fmt.Sprintf("]\n")
	}
	return s
}

// GetSize - return size of matrix
func (m Matrix) GetSize() int {
	return m.size
}
