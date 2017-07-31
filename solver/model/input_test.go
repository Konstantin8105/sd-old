package model_test

import (
	"math"
	"testing"

	"github.com/Konstantin8105/GoFea/input/element"
	"github.com/Konstantin8105/GoFea/input/force"
	"github.com/Konstantin8105/GoFea/input/material"
	"github.com/Konstantin8105/GoFea/input/point"
	"github.com/Konstantin8105/GoFea/input/shape"
	"github.com/Konstantin8105/GoFea/input/support"
	"github.com/Konstantin8105/GoFea/output/displacement"
	"github.com/Konstantin8105/GoFea/output/forceLocal"
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
	// m.AddElement(element.Beam{
	// 	Index:        7,
	// 	PointIndexes: [2]point.Index{4, 2},
	// })
	m.AddElement(element.NewBeam(7, 4, 2))
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
	// m.AddElement(element.Beam{
	// 	Index:        7,
	// 	PointIndexes: [2]point.Index{1, 2},
	// })
	m.AddElement(element.NewBeam(7, 1, 2))
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
	// m.AddElement(element.Beam{
	// 	Index:        7,
	// 	PointIndexes: [2]point.Index{1, 2},
	// })
	m.AddElement(element.NewBeam(7, 1, 2))
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
	// m.AddElement(element.Beam{
	// 	Index:        7,
	// 	PointIndexes: [2]point.Index{1, 2},
	// })
	m.AddElement(element.NewBeam(7, 1, 2))
	m.AddSupport(support.FixedDim2(), 1)
	m.AddShape(shape.Shape{
		A: 300e-6,
	}, []element.Index{7}...)

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
	// m.AddElement(element.Beam{
	// 	Index:        7,
	// 	PointIndexes: [2]point.Index{1, 2},
	// })
	m.AddElement(element.NewBeam(7, 1, 2))
	m.AddSupport(support.FixedDim2(), 1)
	m.AddShape(shape.Shape{
		A: 300e-6,
	}, []element.Index{7, 8, 9, 10}...)

	err := m.Solve()
	if err == nil {
		t.Errorf("Haven't error for model with shape error")
	}
}

func TestErrorTrussInBend(t *testing.T) {
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
	m.AddElement(element.NewBeam(7, 1, 2))
	m.AddSupport(support.FixedDim2(), 1)
	m.AddShape(shape.Shape{
		A: 300e-6,
	}, []element.Index{7}...)
	m.AddNodeForce(1, force.NodeDim2{
		Fy: -80000.0,
	}, []point.Index{2}...)
	m.AddMaterial(material.Linear{
		E:  2e11,
		Ro: 78500,
	}, []element.Index{7}...)
	m.AddTrussProperty(7)

	err := m.Solve()
	if err == nil {
		t.Errorf("Truss on bend, err = %v", err)
	}
}

func TestErrorZeroLength(t *testing.T) {
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

	m.AddElement([]element.Elementer{
		element.NewBeam(7, 4, 2),
		element.NewBeam(8, 4, 1),
		element.NewBeam(9, 4, 3),
	}...)

	// Truss
	m.AddTrussProperty(7, 8, 9)

	// Supports
	m.AddSupport(support.FixedDim2(), 1)
	m.AddSupport(support.FixedDim2(), 2)
	m.AddSupport(support.FixedDim2(), 3)

	// Shapes
	m.AddShape(shape.Shape{
		A: 300e-6,
	}, []element.Index{7, 9}...)

	m.AddShape(shape.Shape{
		A: 300e-6,
	}, []element.Index{8}...)

	// Materials
	m.AddMaterial(material.Linear{
		E:  2e11,
		Ro: 78500,
	}, []element.Index{7, 8, 9}...)

	// Node force
	m.AddNodeForce(1, force.NodeDim2{
		Fy: -80000.0,
	}, []point.Index{4}...)

	err := m.Solve()
	if err == nil {
		t.Errorf("Haven`t checking on zero length of finite element")
	}
}

func TestErrorNoLoads(t *testing.T) {
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

	m.AddElement([]element.Elementer{
		element.NewBeam(7, 4, 2),
		element.NewBeam(8, 4, 1),
		element.NewBeam(9, 4, 3),
	}...)

	// Truss
	m.AddTrussProperty(7, 8, 9)

	// Supports
	m.AddSupport(support.FixedDim2(), 1)
	m.AddSupport(support.FixedDim2(), 2)
	m.AddSupport(support.FixedDim2(), 3)

	// Shapes
	m.AddShape(shape.Shape{
		A: 300e-6,
	}, []element.Index{7, 9}...)

	m.AddShape(shape.Shape{
		A: 300e-6,
	}, []element.Index{8}...)

	// Materials
	m.AddMaterial(material.Linear{
		E:  2e11,
		Ro: 78500,
	}, []element.Index{7, 8, 9}...)

	err := m.Solve()
	if err == nil {
		t.Errorf("Haven`t checking for calculation without loads")
	}
}

func TestErrorWithInfoAboutPoint(t *testing.T) {
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

	m.AddElement([]element.Elementer{
		element.NewBeam(7, 4, 2),
		element.NewBeam(8, 4, 1),
		element.NewBeam(9, 4, 3),
	}...)

	// Truss
	m.AddTrussProperty(7, 8, 9)

	// Supports
	m.AddSupport(support.FixedDim2(), 1)
	m.AddSupport(support.FixedDim2(), 2)
	m.AddSupport(support.FixedDim2(), 3)

	// Shapes
	m.AddShape(shape.Shape{
		A: 300e-6,
	}, []element.Index{7, 9}...)

	m.AddShape(shape.Shape{
		A: 300e-6,
	}, []element.Index{8}...)

	// Materials
	m.AddMaterial(material.Linear{
		E:  2e11,
		Ro: 78500,
	}, []element.Index{7, 8, 9}...)

	// Node force
	m.AddNodeForce(1, force.NodeDim2{
		Fy: -80000.0,
	}, []point.Index{4}...)

	err := m.Solve()
	if err == nil {
		t.Errorf("Haven`t checking for calculation without information about point")
	}
}

func TestErrorCannotCalculate(t *testing.T) {
	var m model.Dim2

	m.AddPoint(point.Dim2{
		Index: 1,
		X:     0.,
		Y:     0.,
	})

	m.AddPoint(point.Dim2{
		Index: 2,
		X:     1.0,
		Y:     0.,
	})

	m.AddElement([]element.Elementer{
		element.NewBeam(7, 1, 2),
	}...)

	// Truss
	m.AddTrussProperty(7)

	// Supports
	m.AddSupport(support.FixedDim2(), 1)

	// Shapes
	m.AddShape(shape.Shape{
		A: 300e-6,
	}, 7)

	// Materials
	m.AddMaterial(material.Linear{
		E:  2e11,
		Ro: 78500,
	}, []element.Index{7}...)

	// Node force
	m.AddNodeForce(1, force.NodeDim2{
		Fy: -80000.0,
	}, []point.Index{2}...)

	err := m.Solve()
	if err == nil {
		t.Errorf("Haven`t checking for wrong calculation (try bend the truss element).\nmodel=%#v", m)
	}
}

func TestErrorGlobalDisplacementForceReaction(t *testing.T) {
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

	m.AddElement([]element.Elementer{
		element.NewBeam(7, 4, 2),
		element.NewBeam(8, 4, 1),
		element.NewBeam(9, 4, 3),
	}...)

	// Truss
	m.AddTrussProperty(7, 8, 9)

	// Supports
	m.AddSupport(support.FixedDim2(), 1)
	m.AddSupport(support.FixedDim2(), 2)
	m.AddSupport(support.FixedDim2(), 3)

	// Shapes
	m.AddShape(shape.Shape{
		A: 300e-6,
	}, []element.Index{7, 9}...)

	m.AddShape(shape.Shape{
		A: 300e-6,
	}, []element.Index{8}...)

	// Materials
	m.AddMaterial(material.Linear{
		E:  2e11,
		Ro: 78500,
	}, []element.Index{7, 8, 9}...)

	// Node force
	m.AddNodeForce(1, force.NodeDim2{
		Fy: -80000.0,
	}, []point.Index{4}...)

	err := m.Solve()
	if err != nil {
		t.Errorf("Cannot solving. error = %v", err)
	}
	// results
	_, err = m.GetGlobalDisplacement(1, point.Index(100))
	if err == nil {
		t.Errorf("Found global displacement with wrong point index")
	}
	_, err = m.GetGlobalDisplacement(10, point.Index(1))
	if err == nil {
		t.Errorf("Found global displacement with wrong force case")
	}

	_, _, err = m.GetLocalForce(1, 100)
	if err == nil {
		t.Errorf("Found local force with wrong beam index")
	}
	_, _, err = m.GetLocalForce(100, 4)
	if err == nil {
		t.Errorf("Found local force with wrong force case")
	}

	_, err = m.GetReaction(1, 100)
	if err == nil {
		t.Errorf("Found reaction with wrong point index")
	}
	_, err = m.GetReaction(100, 1)
	if err == nil {
		t.Errorf("Found global displacement with wrong force case")
	}
}

func trussFrame() (m model.Dim2, err error) {
	m.AddPoint([]point.Dim2{
		{Index: 1, X: 0.0, Y: 0.0},
		{Index: 2, X: 0.0, Y: 1.2},
		{Index: 3, X: 0.4, Y: 0.0},
		{Index: 4, X: 0.4, Y: 0.6},
		{Index: 5, X: 0.8, Y: 0.0},
	}...)

	m.AddElement([]element.Elementer{
		element.NewBeam(1, 1, 2),
		element.NewBeam(2, 1, 3),
		element.NewBeam(3, 1, 4),
		element.NewBeam(4, 2, 4),
		element.NewBeam(5, 3, 4),
		element.NewBeam(6, 3, 5),
		element.NewBeam(7, 4, 5),
	}...)

	m.AddTrussProperty(1, 2, 3, 4, 5, 6, 7)

	m.AddSupport(support.Dim2{Dx: support.Fix, Dy: support.Fix}, 1)
	m.AddSupport(support.Dim2{Dy: support.Fix}, 3)
	m.AddSupport(support.Dim2{Dy: support.Fix}, 5)

	m.AddShape(shape.Shape{A: 40e-4}, []element.Index{1, 5}...)
	m.AddShape(shape.Shape{A: 64e-4}, []element.Index{2, 6}...)
	m.AddShape(shape.Shape{A: 60e-4}, []element.Index{3, 4, 7}...)

	m.AddMaterial(material.Linear{E: 2e11, Ro: 78500}, []element.Index{1, 2, 3, 4, 5, 6, 7}...)

	m.AddNodeForce(1, force.NodeDim2{Fx: -70000.0}, []point.Index{2}...)
	m.AddNodeForce(1, force.NodeDim2{Fx: 42000.0}, []point.Index{4}...)

	err = m.Solve()
	return m, err
}

func isSame(v1, v2 float64) bool {
	if math.Abs(v1) > 1e-8 {
		if math.Abs((v1-v2)/v1) > 0.01 {
			return false
		}
	}
	if math.Abs(v2) > 1e-8 {
		if math.Abs((v1-v2)/v2) > 0.01 {
			return false
		}
	}
	return true
}

func isSameTrussFrames(actual, expected model.Dim2) bool {
	for i := 0; i < 5; i++ {
		var v1, v2 displacement.Dim2
		v1, _ = actual.GetGlobalDisplacement(1, point.Index(i))
		v2, _ = expected.GetGlobalDisplacement(1, point.Index(i))
		if !isSame(v1.Dx, v2.Dx) {
			return false
		}
		if !isSame(v1.Dy, v2.Dy) {
			return false
		}
		if !isSame(v1.Dm, v2.Dm) {
			return false
		}
	}
	for i := 0; i < 7; i++ {
		var v1, v2 forceLocal.Dim2
		v1, _, _ = actual.GetLocalForce(1, element.Index(i))
		v2, _, _ = expected.GetLocalForce(1, element.Index(i))
		if !isSame(v1.Fx, v2.Fx) {
			return false
		}
		if !isSame(v1.Fy, v2.Fy) {
			return false
		}
		if !isSame(v1.M, v2.M) {
			return false
		}
		_, v1, _ = actual.GetLocalForce(1, element.Index(i))
		_, v2, _ = expected.GetLocalForce(1, element.Index(i))
		if !isSame(v1.Fx, v2.Fx) {
			return false
		}
		if !isSame(v1.Fy, v2.Fy) {
			return false
		}
		if !isSame(v1.M, v2.M) {
			return false
		}
	}
	return true
}

func TestErrorDoubleMaterial(t *testing.T) {
	var m model.Dim2

	m.AddPoint([]point.Dim2{
		{Index: 1, X: 0.0, Y: 0.0},
		{Index: 2, X: 0.0, Y: 1.2},
		{Index: 3, X: 0.4, Y: 0.0},
		{Index: 4, X: 0.4, Y: 0.6},
		{Index: 5, X: 0.8, Y: 0.0},
	}...)

	m.AddElement([]element.Elementer{
		element.NewBeam(1, 1, 2),
		element.NewBeam(2, 1, 3),
		element.NewBeam(3, 1, 4),
		element.NewBeam(4, 2, 4),
		element.NewBeam(5, 3, 4),
		element.NewBeam(6, 3, 5),
		element.NewBeam(7, 4, 5),
	}...)

	m.AddTrussProperty(1, 2, 3, 4, 5, 6, 7)

	m.AddSupport(support.Dim2{Dx: support.Fix, Dy: support.Fix}, 1)
	m.AddSupport(support.Dim2{Dy: support.Fix}, 3)
	m.AddSupport(support.Dim2{Dy: support.Fix}, 5)

	m.AddShape(shape.Shape{A: 40e-4}, []element.Index{1, 5}...)
	m.AddShape(shape.Shape{A: 64e-4}, []element.Index{2, 6}...)
	m.AddShape(shape.Shape{A: 60e-4}, []element.Index{3, 4, 7}...)

	m.AddMaterial(material.Linear{E: -1e1, Ro: -8500}, []element.Index{1, 2, 3, 4, 5, 6, 7}...)
	m.AddMaterial(material.Linear{E: 2e11, Ro: 78500}, []element.Index{1, 2, 3, 4, 5, 6, 7}...)

	m.AddNodeForce(1, force.NodeDim2{Fx: -70000.0}, []point.Index{2}...)
	m.AddNodeForce(1, force.NodeDim2{Fx: 42000.0}, []point.Index{4}...)

	err := m.Solve()
	if err == nil {
		t.Errorf("Not correct for double material. error = %v", err)
	}
}

func TestErrorDoubleShape(t *testing.T) {
	var m model.Dim2

	m.AddPoint([]point.Dim2{
		{Index: 1, X: 0.0, Y: 0.0},
		{Index: 2, X: 0.0, Y: 1.2},
		{Index: 3, X: 0.4, Y: 0.0},
		{Index: 4, X: 0.4, Y: 0.6},
		{Index: 5, X: 0.8, Y: 0.0},
	}...)

	m.AddElement([]element.Elementer{
		element.NewBeam(1, 1, 2),
		element.NewBeam(2, 1, 3),
		element.NewBeam(3, 1, 4),
		element.NewBeam(4, 2, 4),
		element.NewBeam(5, 3, 4),
		element.NewBeam(6, 3, 5),
		element.NewBeam(7, 4, 5),
	}...)

	m.AddTrussProperty(1, 2, 3, 4, 5, 6, 7)

	m.AddSupport(support.Dim2{Dx: support.Fix, Dy: support.Fix}, 1)
	m.AddSupport(support.Dim2{Dy: support.Fix}, 3)
	m.AddSupport(support.Dim2{Dy: support.Fix}, 5)

	m.AddShape(shape.Shape{A: 40e-4}, []element.Index{1, 5}...)
	m.AddShape(shape.Shape{A: -64e-4}, []element.Index{2, 6}...)
	m.AddShape(shape.Shape{A: 64e-4}, []element.Index{2, 6}...)
	m.AddShape(shape.Shape{A: 60e-4}, []element.Index{3, 4, 7}...)

	m.AddMaterial(material.Linear{E: 2e11, Ro: 78500}, []element.Index{1, 2, 3, 4, 5, 6, 7}...)

	m.AddNodeForce(1, force.NodeDim2{Fx: -70000.0}, []point.Index{2}...)
	m.AddNodeForce(1, force.NodeDim2{Fx: 42000.0}, []point.Index{4}...)

	err := m.Solve()
	if err == nil {
		t.Errorf("Not correct for double shape. error = %v", err)
	}
}

func TestErrorDoubleTruss(t *testing.T) {
	var m model.Dim2

	m.AddPoint([]point.Dim2{
		{Index: 1, X: 0.0, Y: 0.0},
		{Index: 2, X: 0.0, Y: 1.2},
		{Index: 3, X: 0.4, Y: 0.0},
		{Index: 4, X: 0.4, Y: 0.6},
		{Index: 5, X: 0.8, Y: 0.0},
	}...)

	m.AddElement([]element.Elementer{
		element.NewBeam(1, 1, 2),
		element.NewBeam(2, 1, 3),
		element.NewBeam(3, 1, 4),
		element.NewBeam(4, 2, 4),
		element.NewBeam(5, 3, 4),
		element.NewBeam(6, 3, 5),
		element.NewBeam(7, 4, 5),
	}...)

	m.AddTrussProperty(1, 2, 3, 4, 5, 6, 7)
	m.AddTrussProperty(1, 2, 3, 4, 5, 6, 7)

	m.AddSupport(support.Dim2{Dx: support.Fix, Dy: support.Fix}, 1)
	m.AddSupport(support.Dim2{Dy: support.Fix}, 3)
	m.AddSupport(support.Dim2{Dy: support.Fix}, 5)

	m.AddShape(shape.Shape{A: 40e-4}, []element.Index{1, 5}...)
	m.AddShape(shape.Shape{A: 64e-4}, []element.Index{2, 6}...)
	m.AddShape(shape.Shape{A: 60e-4}, []element.Index{3, 4, 7}...)

	m.AddMaterial(material.Linear{E: 2e11, Ro: 78500}, []element.Index{1, 2, 3, 4, 5, 6, 7}...)

	m.AddNodeForce(1, force.NodeDim2{Fx: -70000.0}, []point.Index{2}...)
	m.AddNodeForce(1, force.NodeDim2{Fx: 42000.0}, []point.Index{4}...)

	err := m.Solve()
	if err == nil {
		t.Errorf("Not correct for double truss. error = %v", err)
	}
}

func TestErrorDoubleSupport(t *testing.T) {
	var m model.Dim2

	m.AddPoint([]point.Dim2{
		{Index: 1, X: 0.0, Y: 0.0},
		{Index: 2, X: 0.0, Y: 1.2},
		{Index: 3, X: 0.4, Y: 0.0},
		{Index: 4, X: 0.4, Y: 0.6},
		{Index: 5, X: 0.8, Y: 0.0},
	}...)

	m.AddElement([]element.Elementer{
		element.NewBeam(1, 1, 2),
		element.NewBeam(2, 1, 3),
		element.NewBeam(3, 1, 4),
		element.NewBeam(4, 2, 4),
		element.NewBeam(5, 3, 4),
		element.NewBeam(6, 3, 5),
		element.NewBeam(7, 4, 5),
	}...)

	m.AddTrussProperty(1, 2, 3, 4, 5, 6, 7)

	m.AddSupport(support.Dim2{Dx: support.Fix, Dy: support.Fix}, 1)
	m.AddSupport(support.Dim2{Dy: support.Fix}, 3)
	m.AddSupport(support.Dim2{Dy: support.Fix}, 3)
	m.AddSupport(support.Dim2{Dy: support.Fix}, 5)

	m.AddShape(shape.Shape{A: 40e-4}, []element.Index{1, 5}...)
	m.AddShape(shape.Shape{A: 64e-4}, []element.Index{2, 6}...)
	m.AddShape(shape.Shape{A: 60e-4}, []element.Index{3, 4, 7}...)

	m.AddMaterial(material.Linear{E: 2e11, Ro: 78500}, []element.Index{1, 2, 3, 4, 5, 6, 7}...)

	m.AddNodeForce(1, force.NodeDim2{Fx: -70000.0}, []point.Index{2}...)
	m.AddNodeForce(1, force.NodeDim2{Fx: 42000.0}, []point.Index{4}...)

	err := m.Solve()
	if err == nil {
		t.Errorf("Not correct for double . error = %v", err)
	}
}

func TestErrorDoubleNodeForce(t *testing.T) {
	var m model.Dim2

	m.AddPoint([]point.Dim2{
		{Index: 1, X: 0.0, Y: 0.0},
		{Index: 2, X: 0.0, Y: 1.2},
		{Index: 3, X: 0.4, Y: 0.0},
		{Index: 4, X: 0.4, Y: 0.6},
		{Index: 5, X: 0.8, Y: 0.0},
	}...)

	m.AddElement([]element.Elementer{
		element.NewBeam(1, 1, 2),
		element.NewBeam(2, 1, 3),
		element.NewBeam(3, 1, 4),
		element.NewBeam(4, 2, 4),
		element.NewBeam(5, 3, 4),
		element.NewBeam(6, 3, 5),
		element.NewBeam(7, 4, 5),
	}...)

	m.AddTrussProperty(1, 2, 3, 4, 5, 6, 7)

	m.AddSupport(support.Dim2{Dx: support.Fix, Dy: support.Fix}, 1)
	m.AddSupport(support.Dim2{Dy: support.Fix}, 3)
	m.AddSupport(support.Dim2{Dy: support.Fix}, 5)

	m.AddShape(shape.Shape{A: 40e-4}, []element.Index{1, 5}...)
	m.AddShape(shape.Shape{A: 64e-4}, []element.Index{2, 6}...)
	m.AddShape(shape.Shape{A: 60e-4}, []element.Index{3, 4, 7}...)

	m.AddMaterial(material.Linear{E: 2e11, Ro: 78500}, []element.Index{1, 2, 3, 4, 5, 6, 7}...)

	m.AddNodeForce(1, force.NodeDim2{Fx: -70000.0}, []point.Index{2}...)
	m.AddNodeForce(1, force.NodeDim2{Fy: -70000.0}, []point.Index{2}...)
	m.AddNodeForce(1, force.NodeDim2{Fx: 42000.0}, []point.Index{4}...)

	err := m.Solve()
	if err == nil {
		t.Errorf("Not correct for double . error = %v", err)
	}
}

func TestErrorDoublePointIndex(t *testing.T) {
	var m model.Dim2

	m.AddPoint([]point.Dim2{
		{Index: 1, X: 10.0, Y: 10.0},
		{Index: 2, X: 0.0, Y: 1.2},
		{Index: 3, X: 0.4, Y: 0.0},
		{Index: 4, X: 0.4, Y: 0.6},
		{Index: 5, X: 0.8, Y: 0.0},
		{Index: 1, X: 0.0, Y: 0.0},
	}...)

	m.AddElement([]element.Elementer{
		element.NewBeam(1, 1, 2),
		element.NewBeam(2, 1, 3),
		element.NewBeam(3, 1, 4),
		element.NewBeam(4, 2, 4),
		element.NewBeam(5, 3, 4),
		element.NewBeam(6, 3, 5),
		element.NewBeam(7, 4, 5),
	}...)

	m.AddTrussProperty(1, 2, 3, 4, 5, 6, 7)

	m.AddSupport(support.Dim2{Dx: support.Fix, Dy: support.Fix}, 1)
	m.AddSupport(support.Dim2{Dy: support.Fix}, 3)
	m.AddSupport(support.Dim2{Dy: support.Fix}, 5)

	m.AddShape(shape.Shape{A: 40e-4}, []element.Index{1, 5}...)
	m.AddShape(shape.Shape{A: 64e-4}, []element.Index{2, 6}...)
	m.AddShape(shape.Shape{A: 60e-4}, []element.Index{3, 4, 7}...)

	m.AddMaterial(material.Linear{E: 2e11, Ro: 78500}, []element.Index{1, 2, 3, 4, 5, 6, 7}...)

	m.AddNodeForce(1, force.NodeDim2{Fx: -70000.0}, []point.Index{2}...)
	m.AddNodeForce(1, force.NodeDim2{Fx: 42000.0}, []point.Index{4}...)

	err := m.Solve()
	if err == nil {
		t.Errorf("Not correct for double . error = %v", err)
	}
}

func TestErrorDoubleElements(t *testing.T) {
	var m model.Dim2

	m.AddPoint([]point.Dim2{
		{Index: 1, X: 0.0, Y: 0.0},
		{Index: 2, X: 0.0, Y: 1.2},
		{Index: 3, X: 0.4, Y: 0.0},
		{Index: 4, X: 0.4, Y: 0.6},
		{Index: 5, X: 0.8, Y: 0.0},
	}...)

	m.AddElement([]element.Elementer{
		element.NewBeam(1, 1, 2),
		element.NewBeam(2, 1, 3),
		element.NewBeam(3, 1, 4),
		element.NewBeam(4, 2, 4),
		element.NewBeam(5, 3, 4),
		element.NewBeam(5, 3, 4),
		element.NewBeam(6, 3, 5),
		element.NewBeam(7, 4, 5),
	}...)

	m.AddTrussProperty(1, 2, 3, 4, 5, 6, 7)

	m.AddSupport(support.Dim2{Dx: support.Fix, Dy: support.Fix}, 1)
	m.AddSupport(support.Dim2{Dy: support.Fix}, 3)
	m.AddSupport(support.Dim2{Dy: support.Fix}, 5)

	m.AddShape(shape.Shape{A: 40e-4}, []element.Index{1, 5}...)
	m.AddShape(shape.Shape{A: 64e-4}, []element.Index{2, 6}...)
	m.AddShape(shape.Shape{A: 60e-4}, []element.Index{3, 4, 7}...)

	m.AddMaterial(material.Linear{E: 2e11, Ro: 78500}, []element.Index{1, 2, 3, 4, 5, 6, 7}...)

	m.AddNodeForce(1, force.NodeDim2{Fx: -70000.0}, []point.Index{2}...)
	m.AddNodeForce(1, force.NodeDim2{Fy: -70000.0}, []point.Index{2}...)
	m.AddNodeForce(1, force.NodeDim2{Fx: 42000.0}, []point.Index{4}...)

	err := m.Solve()
	if err == nil {
		t.Errorf("Not correct for double . error = %v", err)
	}
}

func TestErrorShapeSearch(t *testing.T) {
	var m model.Dim2

	m.AddPoint([]point.Dim2{
		{Index: 1, X: 0.0, Y: 0.0},
		{Index: 2, X: 0.0, Y: 1.2},
		{Index: 3, X: 0.4, Y: 0.0},
		{Index: 4, X: 0.4, Y: 0.6},
		{Index: 5, X: 0.8, Y: 0.0},
	}...)

	m.AddElement([]element.Elementer{
		element.NewBeam(1, 1, 2),
		element.NewBeam(2, 1, 3),
		element.NewBeam(3, 1, 4),
		element.NewBeam(4, 2, 4),
		element.NewBeam(5, 3, 4),
		element.NewBeam(6, 3, 5),
		element.NewBeam(7, 4, 5),
	}...)

	m.AddTrussProperty(1, 2, 3, 4, 5, 6, 7)

	m.AddSupport(support.Dim2{Dx: support.Fix, Dy: support.Fix}, 1)
	m.AddSupport(support.Dim2{Dy: support.Fix}, 3)
	m.AddSupport(support.Dim2{Dy: support.Fix}, 5)

	m.AddShape(shape.Shape{A: 40e-4}, []element.Index{1, 5}...)
	m.AddShape(shape.Shape{A: 64e-4}, []element.Index{2}...) // 6
	m.AddShape(shape.Shape{A: 60e-4}, []element.Index{3, 4, 7}...)

	m.AddMaterial(material.Linear{E: 2e11, Ro: 78500}, []element.Index{1, 2, 3, 4, 5, 6, 7}...)

	m.AddNodeForce(1, force.NodeDim2{Fx: -70000.0}, []point.Index{2}...)
	m.AddNodeForce(1, force.NodeDim2{Fx: 42000.0}, []point.Index{4}...)

	err := m.Solve()
	if err == nil {
		t.Errorf("Not correct for shape not found . error = %v", err)
	}
}

func TestErrorMaterialSearch(t *testing.T) {
	var m model.Dim2

	m.AddPoint([]point.Dim2{
		{Index: 1, X: 0.0, Y: 0.0},
		{Index: 2, X: 0.0, Y: 1.2},
		{Index: 3, X: 0.4, Y: 0.0},
		{Index: 4, X: 0.4, Y: 0.6},
		{Index: 5, X: 0.8, Y: 0.0},
	}...)

	m.AddElement([]element.Elementer{
		element.NewBeam(1, 1, 2),
		element.NewBeam(2, 1, 3),
		element.NewBeam(3, 1, 4),
		element.NewBeam(4, 2, 4),
		element.NewBeam(5, 3, 4),
		element.NewBeam(6, 3, 5),
		element.NewBeam(7, 4, 5),
	}...)

	m.AddTrussProperty(1, 2, 3, 4, 5, 6, 7)

	m.AddSupport(support.Dim2{Dx: support.Fix, Dy: support.Fix}, 1)
	m.AddSupport(support.Dim2{Dy: support.Fix}, 3)
	m.AddSupport(support.Dim2{Dy: support.Fix}, 5)

	m.AddShape(shape.Shape{A: 40e-4}, []element.Index{1, 5}...)
	m.AddShape(shape.Shape{A: 64e-4}, []element.Index{2, 6}...)
	m.AddShape(shape.Shape{A: 60e-4}, []element.Index{3, 4, 7}...)

	m.AddMaterial(material.Linear{E: 2e11, Ro: 78500}, []element.Index{1, 3, 2, 4, 5, 7}...) //6

	m.AddNodeForce(1, force.NodeDim2{Fx: -70000.0}, []point.Index{2}...)
	m.AddNodeForce(1, force.NodeDim2{Fx: 42000.0}, []point.Index{4}...)

	err := m.Solve()
	if err == nil {
		t.Errorf("Not correct for material not found . error = %v", err)
	}
}

func TestErrorCoordinateSearch(t *testing.T) {
	var m model.Dim2

	m.AddPoint([]point.Dim2{
		{Index: 1, X: 0.0, Y: 0.0},
		{Index: 2, X: 0.0, Y: 1.2},
		//	{Index: 3, X: 0.4, Y: 0.0},
		{Index: 4, X: 0.4, Y: 0.6},
		{Index: 5, X: 0.8, Y: 0.0},
	}...)

	m.AddElement([]element.Elementer{
		element.NewBeam(1, 1, 2),
		element.NewBeam(2, 1, 3),
		element.NewBeam(3, 1, 4),
		element.NewBeam(4, 2, 4),
		element.NewBeam(5, 3, 4),
		element.NewBeam(6, 3, 5),
		element.NewBeam(7, 4, 5),
	}...)

	m.AddTrussProperty(1, 2, 3, 4, 5, 6, 7)

	m.AddSupport(support.Dim2{Dx: support.Fix, Dy: support.Fix}, 1)
	m.AddSupport(support.Dim2{Dy: support.Fix}, 3)
	m.AddSupport(support.Dim2{Dy: support.Fix}, 5)

	m.AddShape(shape.Shape{A: 40e-4}, []element.Index{1, 5}...)
	m.AddShape(shape.Shape{A: 64e-4}, []element.Index{2, 6}...)
	m.AddShape(shape.Shape{A: 60e-4}, []element.Index{3, 4, 7}...)

	m.AddMaterial(material.Linear{E: 2e11, Ro: 78500}, []element.Index{1, 3, 2, 4, 6, 5, 7}...)

	m.AddNodeForce(1, force.NodeDim2{Fx: -70000.0}, []point.Index{2}...)
	m.AddNodeForce(1, force.NodeDim2{Fx: 42000.0}, []point.Index{4}...)

	err := m.Solve()
	if err == nil {
		t.Errorf("Not correct for point not found . error = %v", err)
	}
}
