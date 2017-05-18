package model

import matrix "github.com/Konstantin8105/GoFea/Matrix"

type finiteElementer interface {
	getStiffinerK(buffer *matrix.Matrix) error
	getCoordinateTransformation(buffer *matrix.Matrix) error
}
