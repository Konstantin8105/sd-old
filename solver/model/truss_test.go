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
	"github.com/Konstantin8105/GoFea/solver/model"
)

//  *2   *1   *3
//   \   |   /
//    7  8  9
//     \ | /
//      \|/
//       *4
func TestTruss(t *testing.T) {
	var m model.Dim2

	m.AddPoint(point.Dim2{
		Index: 2,
		X:     -0.8660254,
		Y:     0.,
	})

	m.AddPoint(point.Dim2{
		Index: 1,
		X:     0.,
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
		element.NewBeam(8, 4, 1),
		element.NewBeam(9, 4, 3),
		element.NewBeam(7, 4, 2),
	}...)

	// Truss
	m.AddTrussProperty(9, 7, 8)

	// Supports
	m.AddSupport(support.FixedDim2(), 1)
	m.AddSupport(support.FixedDim2(), 3)
	m.AddSupport(support.FixedDim2(), 2)

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
	}, []element.Index{9, 8, 7}...)

	// Node force
	m.AddNodeForce(1, force.NodeDim2{
		Fy: -80000.0,
	}, []point.Index{4}...)

	err := m.Solve()
	if err != nil {
		t.Errorf("Cannot solving. error = %v", err)
	}

	// results

	// displacement : 0.870 mm
	// F7 = F9 = 26098 N
	// F8 = 34797 N
	{
		d, err := m.GetGlobalDisplacement(1, point.Index(4))
		if err != nil {
			t.Errorf("Cannot found global displacement. %v", err)
		}
		de := -0.870e-3 // meter
		if math.Abs((d.Dy-de)/de) > 0.01 {
			t.Errorf("global displacement = %v. Expected displacement = %v", d.Dy, de)
		}
	}
	{
		f7 := -26098.
		b, e, err := m.GetLocalForce(1, element.Index(7))
		if err != nil {
			t.Errorf("Cannot found local force. %v", err)
		}
		if math.Abs((math.Abs(b.Fx)-math.Abs(e.Fx))/b.Fx) > 0.01 {
			t.Errorf("Not symmetrical loads. %v %v", b.Fx, e.Fx)
		}
		if math.Abs((f7-b.Fx)/f7) > 0.01 {
			t.Errorf("axial force for beam 7 is %v. Expected = %v", f7, b.Fx)
		}
	}
	{
		f8 := -34797.
		b, e, err := m.GetLocalForce(1, element.Index(8))
		if err != nil {
			t.Errorf("Cannot found local force. %v", err)
		}
		if math.Abs((math.Abs(b.Fx)-math.Abs(e.Fx))/b.Fx) > 0.01 {
			t.Errorf("Not symmetrical loads. %v %v", b.Fx, e.Fx)
		}
		if math.Abs((f8-b.Fx)/f8) > 0.01 {
			t.Errorf("axial force for beam 8 is %v. Expected = %v", f8, b.Fx)
		}
	}
	{
		f9 := -26098.
		b, e, err := m.GetLocalForce(1, element.Index(7))
		if err != nil {
			t.Errorf("Cannot found local force. %v", err)
		}
		if math.Abs((math.Abs(b.Fx)-math.Abs(e.Fx))/b.Fx) > 0.01 {
			t.Errorf("Not symmetrical loads. %v %v", b.Fx, e.Fx)
		}
		if math.Abs((f9-b.Fx)/f9) > 0.01 {
			t.Errorf("axial force for beam 9 is %v. Expected = %v", f9, b.Fx)
		}
	}
}

// test based on methodic
func TestTrussFrame(t *testing.T) {
	var m model.Dim2

	m.AddPoint([]point.Dim2{
		{
			Index: 1,
			X:     0.0,
			Y:     0.0,
		},
		{
			Index: 2,
			X:     0.0,
			Y:     1.2,
		},
		{
			Index: 3,
			X:     0.4,
			Y:     0.0,
		},
		{
			Index: 4,
			X:     0.4,
			Y:     0.6,
		},
		{
			Index: 5,
			X:     0.8,
			Y:     0.0,
		},
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

	// Truss
	m.AddTrussProperty(1, 2, 3, 4, 5, 6, 7)

	// Supports
	m.AddSupport(support.Dim2{
		Dx: support.Fix,
		Dy: support.Fix,
	}, 1)

	m.AddSupport(support.Dim2{
		Dy: support.Fix,
	}, 3)

	m.AddSupport(support.Dim2{
		Dy: support.Fix,
	}, 5)

	// Shapes
	m.AddShape(shape.Shape{
		A: 40e-4,
	}, []element.Index{1, 5}...)

	m.AddShape(shape.Shape{
		A: 64e-4,
	}, []element.Index{2, 6}...)

	m.AddShape(shape.Shape{
		A: 60e-4,
	}, []element.Index{3, 4, 7}...)

	// Materials
	m.AddMaterial(material.Linear{
		E:  2e11,
		Ro: 78500,
	}, []element.Index{1, 2, 3, 4, 5, 6, 7}...)

	// Node force
	m.AddNodeForce(1, force.NodeDim2{
		Fx: 42000.0,
	}, []point.Index{4}...)

	m.AddNodeForce(1, force.NodeDim2{
		Fx: -70000.0,
	}, []point.Index{2}...)

	err := m.Solve()
	if err != nil {
		t.Errorf("Cannot solving. error = %v", err)
	}

	{
		// displacement for point 2:
		Dx := -0.4610423e-3 // m
		Dy := -0.1575000e-3 // m
		d, err := m.GetGlobalDisplacement(1, point.Index(2))
		if err != nil {
			t.Errorf("Cannot found global displacement. %v", err)
		}
		if math.Abs((d.Dx-Dx)/Dx) > 0.01 {
			t.Errorf("point 1. global displacement by axe X = %v. Expected displacement = %v", d.Dx, Dx)
		}
		if math.Abs((d.Dy-Dy)/Dy) > 0.01 {
			t.Errorf("point 1. global displacement by axe Y = %v. Expected displacement = %v", d.Dy, Dy)
		}
	}
	{
		// displacement for point 4:
		Dx := -0.0380192e-3 // m
		Dy := +0.0333751e-3 // m
		d, err := m.GetGlobalDisplacement(1, point.Index(4))
		if err != nil {
			t.Errorf("Cannot found global displacement. %v", err)
		}
		if math.Abs((d.Dx-Dx)/Dx) > 0.01 {
			t.Errorf("point 4. global displacement by axe X = %v. Expected displacement = %v", d.Dx, Dx)
		}
		if math.Abs((d.Dy-Dy)/Dy) > 0.01 {
			t.Errorf("point 4. global displacement by axe Y = %v. Expected displacement = %v", d.Dy, Dy)
		}
	}
	{
		// local force for beam 2
		FxBegin := 34166.633
		b, _, err := m.GetLocalForce(1, element.Index(2))
		if err != nil {
			t.Errorf("Cannot found local force in beam 2. %v", err)
		}
		if math.Abs((FxBegin-b.Fx)/FxBegin) > 0.01 {
			t.Errorf("axial force for beam 2 is %v. Expected = %v", FxBegin, b.Fx)
		}
	}
	{
		// local force for beam 7
		FxBegin := -61594.72633
		b, _, err := m.GetLocalForce(1, element.Index(7))
		if err != nil {
			t.Errorf("Cannot found local force in beam 2. %v", err)
		}
		if math.Abs((FxBegin-b.Fx)/FxBegin) > 0.01 {
			t.Errorf("axial force for beam 7 is %v. Expected = %v", FxBegin, b.Fx)
		}
	}
	{
		// reaction by X in point 1
		Rx := 28000.0
		r, err := m.GetReaction(1, 1)
		if err != nil {
			t.Errorf("Cannot found reaction in point 1. Error = %v", err)
		}
		if math.Abs((r.Fx-Rx)/Rx) > 0.01 {
			t.Errorf("reaction for point 1 by axe X is %v. Expected = %v", r.Fx, Rx)
		}
	}
}
