package finiteElement

import (
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
func (f *TrussDim2) GetCoordinateTransformation(buffer *linearAlgebra.Matrix) error {
	buffer.SetRectangleSize(2, 4)

	lenght := point.LenghtDim2(f.Points)

	lambdaXX := (f.Points[1].X - f.Points[0].X) / lenght
	lambdaXY := (f.Points[1].Y - f.Points[0].Y) / lenght

	buffer.Set(0, 0, lambdaXX)
	buffer.Set(0, 1, lambdaXY)
	buffer.Set(1, 2, lambdaXX)
	buffer.Set(1, 3, lambdaXY)

	return nil
}

// GetStiffinerK - matrix of stiffiner
func (f *TrussDim2) GetStiffinerK(buffer *linearAlgebra.Matrix) error {
	buffer.SetSize(4)

	lenght := point.LenghtDim2(f.Points)

	EFL := f.Material.E * f.Shape.A / lenght

	buffer.Set(0, 0, EFL)
	buffer.Set(1, 0, -EFL)
	buffer.Set(0, 1, -EFL)
	buffer.Set(1, 1, EFL)

	return nil
}
