package model_test

import (
	"testing"

	"github.com/Konstantin8105/GoFea/input/element"
	"github.com/Konstantin8105/GoFea/input/force"
	"github.com/Konstantin8105/GoFea/input/material"
	"github.com/Konstantin8105/GoFea/input/point"
	"github.com/Konstantin8105/GoFea/input/shape"
	"github.com/Konstantin8105/GoFea/input/support"
	"github.com/Konstantin8105/GoFea/solver/model"
)

// book :
// "Сопротивление материалов"
// Учебник для вузов
// 4-е издание
// год 1979
// страница 139-141
//
//  *2   *1   *3
//   \   |   /
//    7  8  9
//     \ | /
//      \|/
//       *4
func TestTruss(t *testing.T) {
	var m model.Dim2

	m.AddPoint(point.Dim2{
		Index: 1,
		X:     0.,
		Y:     0.,
	})

	m.AddPoint(point.Dim2{
		Index: 2,
		X:     -0.8660254,
		Y:     0.,
	})

	m.AddPoint(point.Dim2{
		Index: 3,
		X:     0.8660254,
		Y:     0.,
	})

	m.AddPoint(point.Dim2{
		Index: 4,
		X:     0.,
		Y:     -1.5,
	})

	// add empty point
	m.AddPoint(point.Dim2{
		Index: 40,
		X:     10.,
		Y:     0.0,
	})

	m.AddElement(element.Beam{
		Index:        7,
		PointIndexes: [2]point.Index{4, 2},
	})

	m.AddElement(element.Beam{
		Index:        8,
		PointIndexes: [2]point.Index{4, 1},
	})

	m.AddElement(element.Beam{
		Index:        9,
		PointIndexes: [2]point.Index{4, 3},
	})

	// Truss
	m.AddTrussProperty(7, 8, 9)

	// Supports
	m.AddSupport(support.FixedDim2(), 1)
	m.AddSupport(support.FixedDim2(), 2)
	m.AddSupport(support.FixedDim2(), 3)

	// Shapes
	m.AddShape(shape.Shape{
		A: 300e-6,
	}, []element.ElementIndex{7, 9}...)

	m.AddShape(shape.Shape{
		A: 300e-6,
	}, []element.ElementIndex{8}...)

	// Materials
	m.AddMaterial(material.Linear{
		E:  2e11,
		Ro: 78500,
	}, []element.ElementIndex{7, 8, 9}...)

	// Node force
	m.AddNodeForce(1, force.NodeDim2{
		Fy: -80000.0,
	}, []point.Index{4}...)

	err := m.Solve()
	if err != nil {
		t.Errorf("Cannot solving. error = %v", err)
	}

	// results

	// displacement : 0.849 mm
	// F7 = F9 = 26097.87104956486
	// F8 = 34797.16132338987
	// Reaction in point 3:
	// Fx = 13048.934 N
	// Fy = 22601.425 N

	//TODO: create test for natural frequency

}

/*
func TestTruss(t *testing.T) {
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
		E:  2e11,
		Ro: 78500,
	}, []element.BeamIndex{1, 2, 3, 4, 5, 6, 7}...)

	// Node force
	model.AddNodeForce(1, force.NodeDim2{
		Fx: -70000.0,
	}, []point.Index{2}...)

	model.AddNodeForce(1, force.NodeDim2{
		Fx: 42000.0,
	}, []point.Index{4}...)

	fmt.Println(model.Solve())

	//TODO: create test for natural frequency

}*/
