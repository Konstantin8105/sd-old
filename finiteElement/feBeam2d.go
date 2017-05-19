package finiteElement

import (
	"github.com/Konstantin8105/GoFea/linearAlgebra"
	"github.com/Konstantin8105/GoFea/material"
	"github.com/Konstantin8105/GoFea/point"
	"github.com/Konstantin8105/GoFea/shape"
)

// BeamDim2 - beam on 2D interpratation
type BeamDim2 struct {
	Material material.Linear
	Shape    shape.Shape
	Points   [2]point.Dim2
}

// GetCoordinateTransformation - record into buffer a matrix of transform from local to global system coordinate
func (f *BeamDim2) GetCoordinateTransformation(buffer *linearAlgebra.Matrix) error {
	panic("TODO")
}

// GetStiffinerK - matrix of stiffiner
func (f *BeamDim2) GetStiffinerK(buffer *linearAlgebra.Matrix) error {
	panic("TODO")
}
