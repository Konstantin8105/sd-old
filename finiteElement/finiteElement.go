package finiteElement

import (
	"github.com/Konstantin8105/GoFea/dof"
	"github.com/Konstantin8105/GoLinAlg/linAlg"
)

var _ FiniteElementer = (*TrussDim2)(nil)

// FiniteElementer - base interface for finite element
type FiniteElementer interface {
	GetCoordinateTransformation(tr *linAlg.Matrix64)
	GetStiffinerK(kr *linAlg.Matrix64)
	GetMassMr(mr *linAlg.Matrix64)
	GetPotentialGr(gr *linAlg.Matrix64, localAxialForce float64)
	GetDoF(degrees *dof.DoF) (axes []dof.AxeNumber)
	//GetStiffinerGlobalK(degree *dof.DoF, info Information) (linAlg.Matrix64, []dof.AxeNumber)
}
