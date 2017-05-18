package finiteElement

import "github.com/Konstantin8105/GoFea/linearAlgebra"

type finiteElementer interface {
	GetStiffinerK(buffer *linearAlgebra.Matrix) error
	GetCoordinateTransformation(buffer *linearAlgebra.Matrix) error
}
