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

func TestErrorEmpty(t *testing.T) {
	var m model.Dim2
	err := m.Solve()
	if err == nil {
		t.Errorf("Haven't error for solving empty model")
	}
}

func TestErrorOneNode(t *testing.T) {
	var m model.Dim2
	m.AddPoint(point.Dim2{
		Index: 1,
		X:     0.,
		Y:     0.,
	})
	err := m.Solve()
	if err == nil {
		t.Errorf("Haven't error for single node model")
	}
}

func TestErrorTwoNode(t *testing.T) {
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
	err := m.Solve()
	if err == nil {
		t.Errorf("Haven't error for 2 node model")
	}
}

func TestErrorBeamAlone(t *testing.T) {
	var m model.Dim2
	m.AddElement(element.Beam{
		Index:        7,
		PointIndexes: [2]point.Index{4, 2},
	})
	err := m.Solve()
	if err == nil {
		t.Errorf("Haven't error for one beam alone")
	}
}

func TestErrorBeamWithoutSupport(t *testing.T) {
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
	m.AddElement(element.Beam{
		Index:        7,
		PointIndexes: [2]point.Index{1, 2},
	})
	err := m.Solve()
	if err == nil {
		t.Errorf("Haven't error for model without support")
	}
}

func TestErrorBeamWithoutLoadAndShape(t *testing.T) {
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
	m.AddElement(element.Beam{
		Index:        7,
		PointIndexes: [2]point.Index{1, 2},
	})
	m.AddSupport(support.FixedDim2(), 1)
	err := m.Solve()
	if err == nil {
		t.Errorf("Haven't error for model without load and shape")
	}
}

func TestErrorBeamWithoutLoad(t *testing.T) {
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
	m.AddElement(element.Beam{
		Index:        7,
		PointIndexes: [2]point.Index{1, 2},
	})
	m.AddSupport(support.FixedDim2(), 1)
	m.AddShape(shape.Shape{
		A: 300e-6,
	}, []element.ElementIndex{7}...)

	err := m.Solve()
	if err == nil {
		t.Errorf("Haven't error for model without load and shape")
	}
}

func TestErrorShape(t *testing.T) {
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
	m.AddElement(element.Beam{
		Index:        7,
		PointIndexes: [2]point.Index{1, 2},
	})
	m.AddSupport(support.FixedDim2(), 1)
	m.AddShape(shape.Shape{
		A: 300e-6,
	}, []element.ElementIndex{7, 8, 9, 10}...)

	err := m.Solve()
	if err == nil {
		t.Errorf("Haven't error for model with shape error")
	}
}

func TestErrorBeamWithLoad(t *testing.T) {
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
	m.AddElement(element.Beam{
		Index:        7,
		PointIndexes: [2]point.Index{1, 2},
	})
	m.AddSupport(support.FixedDim2(), 1)
	m.AddShape(shape.Shape{
		A: 300e-6,
	}, []element.ElementIndex{7}...)
	m.AddNodeForce(1, force.NodeDim2{
		Fy: -80000.0,
	}, []point.Index{2}...)
	m.AddMaterial(material.Linear{
		E:  2e11,
		Ro: 78500,
	}, []element.ElementIndex{7}...)
	m.AddTrussProperty(7)

	err := m.Solve()
	if err != nil {
		t.Errorf("Truss on tension, err = %v", err)
	}
}

func TestErrorZeroLenght(t *testing.T) {
	var m model.Dim2

	m.AddPoint(point.Dim2{
		Index: 1,
		X:     0.,
		Y:     -1.5,
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
	if err == nil {
		t.Errorf("Haven`t checking on zero lenght of finite element")
	}
}
