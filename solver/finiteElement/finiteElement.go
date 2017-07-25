package finiteElement

import (
	"github.com/Konstantin8105/GoFea/solver/dof"
	"github.com/Konstantin8105/GoLinAlg/matrix"
)

var _ FiniteElementer = (*TrussDim2)(nil)

// FiniteElementer - base interface for finite element
type FiniteElementer interface {
	GetCoordinateTransformation(tr *matrix.T64)
	GetStiffinerK(kr *matrix.T64)
	//GetMassMr(mr *matrix.T64)
	//GetPotentialGr(gr *matrix.T64, localAxialForce float64)
	GetDoF(degrees *dof.DoF) (axes []dof.AxeNumber)
	//GetStiffinerGlobalK(degree *dof.DoF, info Information) (matrix.T64, []dof.AxeNumber)
}
