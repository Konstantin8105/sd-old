package finiteElement

import (
	"github.com/Konstantin8105/GoFea/dof"
	"github.com/Konstantin8105/GoFea/linearAlgebra"
	"github.com/Konstantin8105/GoFea/material"
	"github.com/Konstantin8105/GoFea/point"
	"github.com/Konstantin8105/GoFea/shape"
)

// TrussDim2 - truss on 2D interpratation
type TrussDim2 struct {
	Material material.Linear
	Shape    shape.Shape
	Points   [2]point.Dim2
}

// GetCoordinateTransformation - record into buffer a matrix of transform from local to global system coordinate
func (f *TrussDim2) GetCoordinateTransformation(buffer *linearAlgebra.Matrix) {
	buffer.SetRectangleSize(2, 4)

	lenght := point.LenghtDim2(f.Points)

	lambdaXX := (f.Points[1].X - f.Points[0].X) / lenght
	lambdaXY := (f.Points[1].Y - f.Points[0].Y) / lenght

	buffer.Set(0, 0, lambdaXX)
	buffer.Set(0, 1, lambdaXY)
	buffer.Set(1, 2, lambdaXX)
	buffer.Set(1, 3, lambdaXY)
}

// GetStiffinerK - matrix of stiffiner
func (f *TrussDim2) GetStiffinerK(buffer *linearAlgebra.Matrix) {
	buffer.SetSize(2)

	lenght := point.LenghtDim2(f.Points)

	EFL := f.Material.E * f.Shape.A / lenght

	buffer.Set(0, 0, EFL)
	buffer.Set(1, 0, -EFL)
	buffer.Set(0, 1, -EFL)
	buffer.Set(1, 1, EFL)
}

// GetDoF - return numbers for degree of freedom
func (f *TrussDim2) GetDoF(degrees *dof.DoF) (axes []dof.AxeNumber) {
	var Axe [2][]dof.AxeNumber
	Axe[0] = degrees.GetDoF(f.Points[0].Index)
	Axe[1] = degrees.GetDoF(f.Points[1].Index)

	inx := 0
	axes = make([]dof.AxeNumber, 4, 4)
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			axes[inx] = Axe[i][j]
			inx++
		}
	}
	return
}
