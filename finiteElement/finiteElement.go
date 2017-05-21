package finiteElement

import (
	"github.com/Konstantin8105/GoFea/dof"
	"github.com/Konstantin8105/GoFea/linearAlgebra"
)

type finiteElementer interface {
	GetStiffinerK(buffer *linearAlgebra.Matrix)
	GetCoordinateTransformation(buffer *linearAlgebra.Matrix)
	GetDoF(degrees *dof.DoF) []dof.AxeNumber
}
