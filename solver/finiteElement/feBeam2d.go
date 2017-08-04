package finiteElement

import (
	"math"

	"github.com/Konstantin8105/GoFea/input/material"
	"github.com/Konstantin8105/GoFea/input/point"
	"github.com/Konstantin8105/GoFea/input/shape"
	"github.com/Konstantin8105/GoFea/solver/dof"
	"github.com/Konstantin8105/GoLinAlg/matrix"
)

// BeamDim2 - truss on 2D interpretation
type BeamDim2 struct {
	Material material.Linear
	Shape    shape.Shape
	Points   [2]point.Dim2
}

// GetCoordinateTransformation - matrix of transform between local and global system coordinate
func (f *BeamDim2) GetCoordinateTransformation(tr *matrix.T64) {
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
func (f *BeamDim2) GetStiffinerK(kr *matrix.T64) {
	length := point.LengthDim2(f.Points)

	kr.SetNewSize(6, 6)

	EFL := f.Material.E * f.Shape.A / length
	kr.Set(0, 0, EFL)
	kr.Set(0, 3, -EFL)
	kr.Set(3, 0, -EFL)
	kr.Set(3, 3, EFL)

	EJL3 := 12.0 * f.Material.E * f.Shape.Izz / math.Pow(length, 3.0)
	kr.Set(1, 1, EJL3)
	kr.Set(4, 4, EJL3)
	kr.Set(1, 4, -EJL3)
	kr.Set(4, 1, -EJL3)

	EJL2 := 6.0 * f.Material.E * f.Shape.Izz / (length * length)
	kr.Set(1, 2, EJL2)
	kr.Set(2, 1, EJL2)
	kr.Set(1, 5, EJL2)
	kr.Set(5, 1, EJL2)
	kr.Set(2, 4, -EJL2)
	kr.Set(4, 2, -EJL2)
	kr.Set(4, 5, -EJL2)
	kr.Set(5, 4, -EJL2)

	EJL := 2.0 * f.Material.E * f.Shape.Izz / (length)
	kr.Set(2, 5, EJL)
	kr.Set(5, 2, EJL)

	EJL = 2 * EJL
	kr.Set(2, 2, EJL)
	kr.Set(5, 5, EJL)
}

// GetMassMr - matrix mass of finite element
func (f *BeamDim2) GetMassMr(mr *matrix.T64) {
	mu := f.Shape.A * f.Material.Ro
	length := point.LengthDim2(f.Points)
	mul3 := length / 3.0 * mu
	mul6 := length / 6.0 * mu

	mr.SetNewSize(6, 6)
	mr.Set(0, 0, mul3)
	mr.Set(0, 3, mul6)
	mr.Set(3, 0, mul6)
	mr.Set(3, 3, mul3)

	{
		v := mu * 13.0 * length / 35.
		mr.Set(1, 1, v)
		mr.Set(4, 4, v)
	}
	{
		v := mu * 11.0 * length * length / 210.0
		mr.Set(1, 2, v)
		mr.Set(2, 1, v)
		mr.Set(4, 5, -v)
		mr.Set(5, 4, -v)
	}
	{
		v := mu * length * length / 105.
		mr.Set(2, 2, v)
	}
	{
		v := mu * math.Pow(length, 3.0)
		mr.Set(5, 5, v)
	}
	{
		v := mu * 9.0 * length / 70.
		mr.Set(1, 4, v)
		mr.Set(4, 1, v)
	}
	{
		v := -mu * math.Pow(length, 3.0) / 140.
		mr.Set(2, 5, v)
		mr.Set(5, 2, v)
	}
	{
		v := mu * 13.0 * length * length / 420.0
		mr.Set(2, 4, v)
		mr.Set(4, 2, v)
		mr.Set(1, 5, -v)
		mr.Set(5, 1, -v)
	}
}

// GetPotentialGr - matrix potential loads for linear buckling
func (f *BeamDim2) GetPotentialGr(gr *matrix.T64, localAxialForce float64) {
	length := point.LengthDim2(f.Points)

	gr.SetNewSize(6, 6)
	{
		NL := 6.0 / 5.0 * localAxialForce / length
		gr.Set(1, 1, NL)
		gr.Set(1, 4, -NL)
		gr.Set(4, 1, -NL)
		gr.Set(4, 4, NL)
	}
	{
		v := 0.1 * localAxialForce
		gr.Set(1, 2, v)
		gr.Set(2, 1, v)
		gr.Set(1, 5, v)
		gr.Set(5, 1, v)
	}
	{
		v := -0.1 * localAxialForce
		gr.Set(2, 4, v)
		gr.Set(4, 2, v)
		gr.Set(4, 5, v)
		gr.Set(5, 4, v)
	}
	{
		v := localAxialForce * 2.0 / 15. * length
		gr.Set(2, 2, v)
		gr.Set(5, 5, v)
	}
	{
		v := -localAxialForce * length / 30.0
		gr.Set(2, 5, v)
		gr.Set(5, 2, v)
	}
}

// GetDoF - return numbers for degree of freedom in global system
// coordinate
func (f *BeamDim2) GetDoF(degrees *dof.DoF) (axes []dof.AxeNumber) {
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
