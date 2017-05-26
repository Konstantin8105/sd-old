package finiteElement

import (
	"github.com/Konstantin8105/GoFea/dof"
	"github.com/Konstantin8105/GoLinAlg/linAlg"
)

// FiniteElementer - base interface for finite element
type FiniteElementer interface {
	GetStiffinerK(buffer *linAlg.Matrix64)
	GetCoordinateTransformation(buffer *linAlg.Matrix64)
	GetDoF(degrees *dof.DoF) []dof.AxeNumber
}
