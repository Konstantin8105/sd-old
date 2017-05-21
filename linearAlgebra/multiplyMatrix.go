package linearAlgebra

// MultiplyTtKT - multiply matrix
// formula: T(transponse) * M * T
func (m *Matrix) MultiplyTtKT(t Matrix, buffer *Matrix) Matrix {
	if t.sizeI != m.sizeI {
		panic("Not correct algoritm")
	}

	buffer.SetRectangleSize(t.sizeJ, m.sizeJ)

	for i := 0; i < buffer.sizeI; i++ {
		for j := 0; j < buffer.sizeJ; j++ {
			sum := 0.0
			for k := 0; k < t.sizeI; k++ {
				sum += t.Get(k, i) * m.Get(k, j)
			}
			buffer.Set(i, j, sum)
		}
	}

	result := NewRectangleMatrix(buffer.sizeI, t.sizeJ)
	for i := 0; i < result.sizeI; i++ {
		for j := 0; j < result.sizeJ; j++ {
			sum := 0.0
			for k := 0; k < buffer.sizeJ; k++ {
				sum += buffer.Get(i, k) * t.Get(k, j)
			}
			result.Set(i, j, sum)
		}
	}

	if result.sizeI == result.sizeJ {
		result.typeM = Square
	}
	return result
}
