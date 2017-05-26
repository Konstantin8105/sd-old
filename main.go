package main

import (
	"fmt"

	"github.com/Konstantin8105/GoFea/element"
	"github.com/Konstantin8105/GoFea/force"
	"github.com/Konstantin8105/GoFea/material"
	"github.com/Konstantin8105/GoFea/model"
	"github.com/Konstantin8105/GoFea/point"
	"github.com/Konstantin8105/GoFea/shape"
	"github.com/Konstantin8105/GoFea/support"
)

func main() {
	truss()
}

func truss() {

	var model model.Dim2

	model.AddPoint([]point.Dim2{
		point.Dim2{
			Index: 1,
			X:     0.0,
			Y:     0.0,
		},
		point.Dim2{
			Index: 2,
			X:     0.0,
			Y:     1.2,
		},
		point.Dim2{
			Index: 3,
			X:     0.4,
			Y:     0.0,
		},
		point.Dim2{
			Index: 4,
			X:     0.4,
			Y:     0.6,
		},
		point.Dim2{
			Index: 5,
			X:     0.8,
			Y:     0.0,
		},
	}...)

	model.AddBeam([]element.Beam{
		element.Beam{
			Index:        1,
			PointIndexes: [2]point.Index{1, 2},
		},
		element.Beam{
			Index:        2,
			PointIndexes: [2]point.Index{1, 3},
		},
		element.Beam{
			Index:        3,
			PointIndexes: [2]point.Index{1, 4},
		},
		element.Beam{
			Index:        4,
			PointIndexes: [2]point.Index{2, 4},
		},
		element.Beam{
			Index:        5,
			PointIndexes: [2]point.Index{3, 4},
		},
		element.Beam{
			Index:        6,
			PointIndexes: [2]point.Index{3, 5},
		},
		element.Beam{
			Index:        7,
			PointIndexes: [2]point.Index{4, 5},
		},
	}...)

	// Truss
	model.AddTrussProperty(1, 2, 3, 4, 5, 6, 7)

	// Supports
	model.AddSupport(support.Dim2{
		Dx: support.Fix,
		Dy: support.Fix,
	}, 1)

	model.AddSupport(support.Dim2{
		Dy: support.Fix,
	}, 3)

	model.AddSupport(support.Dim2{
		Dy: support.Fix,
	}, 5)

	// Shapes
	model.AddShape(shape.Shape{
		A: 40e-4,
	}, []element.BeamIndex{1, 5}...)

	model.AddShape(shape.Shape{
		A: 64e-4,
	}, []element.BeamIndex{2, 6}...)

	model.AddShape(shape.Shape{
		A: 60e-4,
	}, []element.BeamIndex{3, 4, 7}...)

	// Materials
	model.AddMaterial(material.Linear{
		E: 2e11,
	}, []element.BeamIndex{1, 2, 3, 4, 5, 6, 7}...)

	// Node force
	model.AddNodeForce(1, force.NodeDim2{
		Fx: -70000.0,
	}, []point.Index{2}...)

	model.AddNodeForce(1, force.NodeDim2{
		Fx: 42000.0,
	}, []point.Index{4}...)

	fmt.Println(model.Solve())
}
