package finiteElement

import (
	"github.com/Konstantin8105/GoFea/linearAlgebra"
	"github.com/Konstantin8105/GoFea/material"
	"github.com/Konstantin8105/GoFea/point"
	"github.com/Konstantin8105/GoFea/shape"
)

// BeamDim3 - beam on 3D interpratation
type BeamDim3 struct {
	Material material.Linear
	Shape    shape.Shape
	Points   [2]point.Dim3
}

// GetCoordinateTransformation - record into buffer a matrix of transform from local to global system coordinate
func (f BeamDim3) GetCoordinateTransformation(buffer *linearAlgebra.Matrix) (err error) {
	const (
		size = 12
	)
	buffer.SetSize(size)

	//TODO: add algoritm
	panic("Add algoritm")

}

// GetStiffinerK - matrix of stiffiner
func (f BeamDim3) GetStiffinerK(buffer *linearAlgebra.Matrix) (err error) {
	const (
		size = 12
	)
	buffer.SetSize(size)

	lenght := point.LenghtDim3(f.Points)

	// add only in lower triangle of matrix
	{
		EFL := f.Material.E * f.Shape.A / lenght
		buffer.Set(0, 0, EFL)
		buffer.Set(6, 0, -EFL)
		buffer.Set(6, 6, EFL)
	}
	{
		EJzL := 12.0 * f.Material.E * f.Shape.Izz / (lenght * lenght * lenght)
		buffer.Set(1, 1, EJzL)
		buffer.Set(7, 1, -EJzL)
		buffer.Set(7, 7, EJzL)
	}
	{
		EJyL := 12.0 * f.Material.E * f.Shape.Iyy / (lenght * lenght * lenght)
		buffer.Set(2, 2, EJyL)
		buffer.Set(8, 2, -EJyL)
		buffer.Set(8, 8, EJyL)
	}
	{
		GJxL := f.Material.G * f.Shape.Jxx / lenght
		buffer.Set(3, 3, GJxL)
		buffer.Set(9, 3, -GJxL)
		buffer.Set(9, 9, GJxL)
	}
	{
		EJyL := 4.0 * f.Material.E * f.Shape.Iyy / lenght
		buffer.Set(4, 4, EJyL)
		buffer.Set(10, 10, EJyL)
	}
	{
		EJzL := 4.0 * f.Material.E * f.Shape.Izz / lenght
		buffer.Set(5, 5, EJzL)
		buffer.Set(11, 11, EJzL)
	}
	{
		EJyL := 6.0 * f.Material.E * f.Shape.Iyy / (lenght * lenght)
		buffer.Set(4, 2, -EJyL)
		buffer.Set(10, 2, -EJyL)
		buffer.Set(10, 8, EJyL)
		buffer.Set(8, 4, EJyL)
	}
	{
		EJzL := 6.0 * f.Material.E * f.Shape.Izz / (lenght * lenght)
		buffer.Set(5, 1, EJzL)
		buffer.Set(11, 1, EJzL)
		buffer.Set(11, 7, -EJzL)
		buffer.Set(7, 5, -EJzL)
	}
	{
		buffer.Set(10, 4, 2.0*f.Material.E*f.Shape.Iyy/lenght)
	}
	{
		buffer.Set(11, 5, 2.0*f.Material.E*f.Shape.Izz/lenght)
	}

	// repeat to upper triangle of matrix
	for i := 0; i < size; i++ {
		for j := 0; j < i; j++ {
			buffer.Set(j, i, (*buffer).Get(i, j))
		}
	}
	return nil
}
