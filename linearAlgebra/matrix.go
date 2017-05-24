package linearAlgebra

import "fmt"

// Type - type of matrix shape
type Type int

// Types of matrix
const (
	Square Type = iota
	Rectangle
)

// Matrix - matrix data structure
type Matrix struct {
	sizeI, sizeJ int
	capacity     int
	values       [][]float64
	typeM        Type
}

// NewSquareMatrix - constructor for square matrix
func NewSquareMatrix(size int) (m Matrix) {
	//m = *new(Matrix) // ?? no need
	m.typeM = Square
	m.SetSize(size)
	return
}

// GetSize - return size of square matrix
func (m *Matrix) GetSize() int {
	if m.typeM != Square {
		panic("Not correct type of matrix")
	}
	if m.sizeI != m.sizeJ {
		panic("Not correct size")
	}
	return m.sizeI
}

// NewRectangleMatrix - create new rectangle matrix
func NewRectangleMatrix(si, sj int) (m Matrix) {
	m = *new(Matrix)
	m.typeM = Rectangle
	m.SetRectangleSize(si, sj)
	return
}

// SetSize - change to square matrix size
func (m *Matrix) SetSize(s int) {
	m.typeM = Square
	m.sizeI = s
	m.sizeJ = s
	if m.capacity < s {
		m.values = make([][]float64, s, s)
		m.capacity = s
		for i := 0; i < s; i++ {
			m.values[i] = make([]float64, s)
		}
		return
	}
	// initialize the matrix
	for i := 0; i < m.sizeI; i++ {
		for j := 0; j < m.sizeJ; j++ {
			m.values[i][j] = 0.0
		}
	}
}

// SetRectangleSize - change the matrix size
func (m *Matrix) SetRectangleSize(si, sj int) {
	if si > sj {
		m.SetSize(si)
	} else {
		m.SetSize(sj)
	}
	m.typeM = Rectangle
	m.sizeI = si
	m.sizeJ = sj
}

// Set - change value of matrix
func (m *Matrix) Set(i int, j int, value float64) {
	if i < 0 || i >= m.sizeI || j < 0 || j >= m.sizeJ {
		panic(fmt.Errorf("Cannot index is outside the matrix - [%v,%v]\nMatrix - %#v", i, j, m))
	}
	m.values[i][j] = value
}

// Get - return value in matrix
func (m *Matrix) Get(i int, j int) float64 {
	if i < 0 || i >= m.sizeI || j < 0 || j >= m.sizeJ {
		panic(fmt.Errorf("Cannot index is outside the matrix - [%v,%v]\nMatrix - %#v", i, j, m))
	}
	return m.values[i][j]
}

func (m Matrix) String() (s string) {
	for i := 0; i < m.sizeI; i++ {
		s += fmt.Sprintf("[")
		for j := 0; j < m.sizeJ; j++ {
			s += fmt.Sprintf("%+.2E", m.values[i][j])
			if j != m.sizeJ-1 {
				s += fmt.Sprintf(",")
			}
		}
		s += fmt.Sprintf("]\n")
	}
	return s
}
