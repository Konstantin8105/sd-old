package model

import (
	"fmt"
	"math"

	coordinate "github.com/Konstantin8105/GoFea/Coordinate"
	material "github.com/Konstantin8105/GoFea/Material"
	matrix "github.com/Konstantin8105/GoFea/Matrix"
	shape "github.com/Konstantin8105/GoFea/Shape"
)

// FiniteElement3Dbeam - beam on 3D interpratation
type FiniteElement3Dbeam struct {
	material    material.Material
	shape       shape.Shape
	coordinates [2]coordinate.Coordinate
}

func (f FiniteElement3Dbeam) getCoordinateTransformation(buffer *matrix.Matrix) (err error) {
	const (
		size = 12
	)
	if buffer.Type() != matrix.Square {
		return fmt.Errorf("Cannot insert to not square buffer")
	}
	buffer.SetSize(size)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			buffer.Set(i, j, 0.0)
		}
	}

	return nil
}

// StiffinerK - matrix of stiffiner
func (f FiniteElement3Dbeam) getStiffinerK(buffer *matrix.Matrix) (err error) {
	const (
		size = 12
	)
	if buffer.Type() != matrix.Square {
		return fmt.Errorf("Cannot insert to not square buffer")
	}
	buffer.SetSize(size)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			buffer.Set(i, j, 0.0)
		}
	}

	lenght := math.Sqrt(math.Pow(f.coordinates[0].X-f.coordinates[1].X, 2.0) + math.Pow(f.coordinates[0].Y-f.coordinates[1].Y, 2.0) + math.Pow(f.coordinates[0].Z-f.coordinates[1].Z, 2.0))

	// add only in lower triangle of matrix
	{
		EFL := f.material.E * f.shape.A / lenght
		buffer.Set(0, 0, EFL)
		buffer.Set(6, 0, -EFL)
		buffer.Set(6, 6, EFL)
	}
	{
		EJzL := 12.0 * f.material.E * f.shape.Izz / (lenght * lenght * lenght)
		buffer.Set(1, 1, EJzL)
		buffer.Set(7, 1, -EJzL)
		buffer.Set(7, 7, EJzL)
	}
	{
		EJyL := 12.0 * f.material.E * f.shape.Iyy / (lenght * lenght * lenght)
		buffer.Set(2, 2, EJyL)
		buffer.Set(8, 2, -EJyL)
		buffer.Set(8, 8, EJyL)
	}
	{
		GJxL := f.material.G * f.shape.Jxx / lenght
		buffer.Set(3, 3, GJxL)
		buffer.Set(9, 3, -GJxL)
		buffer.Set(9, 9, GJxL)
	}
	{
		EJyL := 4.0 * f.material.E * f.shape.Iyy / lenght
		buffer.Set(4, 4, EJyL)
		buffer.Set(10, 10, EJyL)
	}
	{
		EJzL := 4.0 * f.material.E * f.shape.Izz / lenght
		buffer.Set(5, 5, EJzL)
		buffer.Set(11, 11, EJzL)
	}
	{
		EJyL := 6.0 * f.material.E * f.shape.Iyy / (lenght * lenght)
		buffer.Set(4, 2, -EJyL)
		buffer.Set(10, 2, -EJyL)
		buffer.Set(10, 8, EJyL)
		buffer.Set(8, 4, EJyL)
	}
	{
		EJzL := 6.0 * f.material.E * f.shape.Izz / (lenght * lenght)
		buffer.Set(5, 1, EJzL)
		buffer.Set(11, 1, EJzL)
		buffer.Set(11, 7, -EJzL)
		buffer.Set(7, 5, -EJzL)
	}
	{
		buffer.Set(10, 4, 2.0*f.material.E*f.shape.Iyy/lenght)
	}
	{
		buffer.Set(11, 5, 2.0*f.material.E*f.shape.Izz/lenght)
	}

	// repeat to upper triangle of matrix
	for i := 0; i < size; i++ {
		for j := 0; j < i; j++ {
			buffer.Set(j, i, (*buffer).Get(i, j))
		}
	}
	return nil
}
