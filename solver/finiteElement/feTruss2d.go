package finiteElement

import (
	"github.com/Konstantin8105/GoFea/input/material"
	"github.com/Konstantin8105/GoFea/input/point"
	"github.com/Konstantin8105/GoFea/input/shape"
	"github.com/Konstantin8105/GoFea/solver/dof"
	"github.com/Konstantin8105/GoLinAlg/matrix"
)

// TrussDim2 - truss on 2D interpretation
type TrussDim2 struct {
	Material material.Linear
	Shape    shape.Shape
	Points   [2]point.Dim2
}

// GetCoordinateTransformation - matrix of transform between local and global system coordinate
func (f *TrussDim2) GetCoordinateTransformation(tr *matrix.T64) {
	length := point.LengthDim2(f.Points)
	lambdaXX := (f.Points[1].X - f.Points[0].X) / length
	lambdaXY := (f.Points[1].Y - f.Points[0].Y) / length

	tr.SetNewSize(6, 6)
	tr.Set(0, 0, lambdaXX)
	tr.Set(0, 1, lambdaXY)
	tr.Set(1, 0, -lambdaXY)
	tr.Set(1, 1, lambdaXX)
	tr.Set(2, 2, 1.0)

	tr.Set(3, 3, lambdaXX)
	tr.Set(3, 4, lambdaXY)
	tr.Set(4, 3, -lambdaXY)
	tr.Set(4, 4, lambdaXX)
	tr.Set(5, 5, 1.0)

}

// GetStiffinerK - matrix of stiffiner
func (f *TrussDim2) GetStiffinerK(kr *matrix.T64) {
	length := point.LengthDim2(f.Points)
	EFL := f.Material.E * f.Shape.A / length

	kr.SetNewSize(6, 6)
	kr.Set(0, 0, EFL)
	kr.Set(0, 3, -EFL)
	kr.Set(3, 0, -EFL)
	kr.Set(3, 3, EFL)
}

/*
// GetMassMr - matrix mass of finite element
func (f *TrussDim2) GetMassMr(mr *matrix.T64) {
	mu := f.Shape.A * f.Material.Ro
	length := point.LenghtDim2(f.Points)
	mul3 := length / 3.0 * mu
	mul6 := length / 6.0 * mu

	mr.SetNewSize(6, 6)
	mr.Set(0, 0, mul3)
	mr.Set(0, 3, mul6)
	mr.Set(3, 0, mul6)
	mr.Set(3, 3, mul3)
}
*/

/*
// GetPotentialGr - matrix potential loads for linear buckling
func (f *TrussDim2) GetPotentialGr(gr *matrix.T64, localAxialForce float64) {
	length := point.LenghtDim2(f.Points)

	NL := localAxialForce / length
	// TODO check somewhere length cannot by zero

	gr.SetNewSize(6, 6)
	gr.Set(1, 1, NL)
	gr.Set(1, 4, -NL)
	gr.Set(4, 1, -NL)
	gr.Set(4, 4, NL)
}
*/

// GetDoF - return numbers for degree of freedom in global system
// coordinate
func (f *TrussDim2) GetDoF(degrees *dof.DoF) (axes []dof.AxeNumber) {
	var Axe [2][]dof.AxeNumber
	Axe[0] = degrees.GetDoF(f.Points[0].Index)
	Axe[1] = degrees.GetDoF(f.Points[1].Index)

	inx := 0
	axes = make([]dof.AxeNumber, 6)
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			axes[inx] = Axe[i][j]
			inx++
		}
	}
	return
}
