package model_test

// DO NOT EDIT!!
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
func TestGenTrussSuccessFull(t *testing.T) {
	var m model.Dim2
	m.AddPoint(point.Dim2{Index: 2, X: -0.8660254, Y: 0.})
	m.AddPoint(point.Dim2{Index: 1, X: 0., Y: 0.})
	m.AddPoint(point.Dim2{Index: 3, X: 0.8660254, Y: 0.})
	m.AddPoint(point.Dim2{Index: 4, X: 0., Y: -1.5})
	m.AddPoint(point.Dim2{Index: 40, X: 10., Y: 0.0})
	m.AddElement([]element.Elementer{
		element.NewBeam(8, 4, 1),
		element.NewBeam(9, 4, 3),
	}...)
	m.AddElement([]element.Elementer{
		element.NewBeam(7, 4, 2),
	}...)
	m.AddTrussProperty(7)
	m.AddTrussProperty(9, 8)
	m.AddSupport(support.FixedDim2(), 1)
	m.AddSupport(support.FixedDim2(), 3)
	m.AddSupport(support.FixedDim2(), 2)
	m.AddShape(shape.Shape{
		A: 300e-6,
	}, []element.Index{7, 9}...)
	m.AddShape(shape.Shape{
		A: 300e-6,
	}, []element.Index{8}...)
	m.AddMaterial(material.Linear{
		E:  2e11,
		Ro: 78500,
	}, []element.Index{9, 7}...)
	m.AddMaterial(material.Linear{
		E:  2e11,
		Ro: 78500,
	}, []element.Index{8}...)
	m.AddNodeForce(1, force.NodeDim2{
		Fy: -80000.0,
	}, []point.Index{4}...)
	m.AddNaturalFrequency(2)
	m.AddNodeForce(2, force.NodeDim2{
		Fx: 10000.0,
		Fy: 10000.0,
	}, []point.Index{4}...)
	err := m.Solve()
	if err != nil {
		t.Errorf("Cannot solving. error = %v", err)
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
}
func TestGenTrussSuccessInvertFull(t *testing.T) {
	var m model.Dim2
	m.AddNodeForce(1, force.NodeDim2{
		Fy: -80000.0,
	}, []point.Index{4}...)
	m.AddMaterial(material.Linear{
		E:  2e11,
		Ro: 78500,
	}, []element.Index{8}...)
	m.AddMaterial(material.Linear{
		E:  2e11,
		Ro: 78500,
	}, []element.Index{9, 7}...)
	m.AddShape(shape.Shape{
		A: 300e-6,
	}, []element.Index{8}...)
	m.AddShape(shape.Shape{
		A: 300e-6,
	}, []element.Index{7, 9}...)
	m.AddSupport(support.FixedDim2(), 3)
	m.AddSupport(support.FixedDim2(), 2)
	m.AddSupport(support.FixedDim2(), 1)
	m.AddTrussProperty(9, 8)
	m.AddTrussProperty(7)
	m.AddElement([]element.Elementer{
		element.NewBeam(7, 4, 2),
	}...)
	m.AddElement([]element.Elementer{
		element.NewBeam(8, 4, 1),
		element.NewBeam(9, 4, 3),
	}...)
	m.AddPoint(point.Dim2{Index: 3, X: 0.8660254, Y: 0.})
	m.AddPoint(point.Dim2{Index: 4, X: 0., Y: -1.5})
	m.AddPoint(point.Dim2{Index: 40, X: 10., Y: 0.0})
	m.AddPoint(point.Dim2{Index: 2, X: -0.8660254, Y: 0.})
	m.AddPoint(point.Dim2{Index: 1, X: 0., Y: 0.})
	m.AddNaturalFrequency(2)
	m.AddNodeForce(2, force.NodeDim2{
		Fx: 10000.0,
		Fy: 10000.0,
	}, []point.Index{4}...)
	err := m.Solve()
	if err != nil {
		t.Errorf("Cannot solving. error = %v", err)
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
}
func TestGenTrussWothoutOneForFail0(t *testing.T) {
	var m model.Dim2
	m.AddPoint(point.Dim2{Index: 3, X: 0.8660254, Y: 0.})
	m.AddPoint(point.Dim2{Index: 4, X: 0., Y: -1.5})
	m.AddPoint(point.Dim2{Index: 40, X: 10., Y: 0.0})
	m.AddElement([]element.Elementer{
		element.NewBeam(8, 4, 1),
		element.NewBeam(9, 4, 3),
	}...)
	m.AddElement([]element.Elementer{
		element.NewBeam(7, 4, 2),
	}...)
	m.AddTrussProperty(7)
	m.AddTrussProperty(9, 8)
	m.AddSupport(support.FixedDim2(), 1)
	m.AddSupport(support.FixedDim2(), 3)
	m.AddSupport(support.FixedDim2(), 2)
	m.AddShape(shape.Shape{
		A: 300e-6,
	}, []element.Index{7, 9}...)
	m.AddShape(shape.Shape{
		A: 300e-6,
	}, []element.Index{8}...)
	m.AddMaterial(material.Linear{
		E:  2e11,
		Ro: 78500,
	}, []element.Index{9, 7}...)
	m.AddMaterial(material.Linear{
		E:  2e11,
		Ro: 78500,
	}, []element.Index{8}...)
	m.AddNodeForce(1, force.NodeDim2{
		Fy: -80000.0,
	}, []point.Index{4}...)
	m.AddNaturalFrequency(2)
	m.AddNodeForce(2, force.NodeDim2{
		Fx: 10000.0,
		Fy: 10000.0,
	}, []point.Index{4}...)
	err := m.Solve()
	if err == nil {
		f7 := -26098.
		b, _, err := m.GetLocalForce(1, element.Index(7))
		if err != nil {
			return
		}
		if math.Abs((f7-b.Fx)/f7) > 0.01 {
			return
		} else {
			t.Errorf("axial force for beam 7 is %v cannot be equal without some data. Expected = %v", f7, b.Fx)
		}
	}
}
func TestGenTrussWothoutOneForFail1(t *testing.T) {
	var m model.Dim2
	m.AddPoint(point.Dim2{Index: 2, X: -0.8660254, Y: 0.})
	m.AddPoint(point.Dim2{Index: 1, X: 0., Y: 0.})
	m.AddElement([]element.Elementer{
		element.NewBeam(8, 4, 1),
		element.NewBeam(9, 4, 3),
	}...)
	m.AddElement([]element.Elementer{
		element.NewBeam(7, 4, 2),
	}...)
	m.AddTrussProperty(7)
	m.AddTrussProperty(9, 8)
	m.AddSupport(support.FixedDim2(), 1)
	m.AddSupport(support.FixedDim2(), 3)
	m.AddSupport(support.FixedDim2(), 2)
	m.AddShape(shape.Shape{
		A: 300e-6,
	}, []element.Index{7, 9}...)
	m.AddShape(shape.Shape{
		A: 300e-6,
	}, []element.Index{8}...)
	m.AddMaterial(material.Linear{
		E:  2e11,
		Ro: 78500,
	}, []element.Index{9, 7}...)
	m.AddMaterial(material.Linear{
		E:  2e11,
		Ro: 78500,
	}, []element.Index{8}...)
	m.AddNodeForce(1, force.NodeDim2{
		Fy: -80000.0,
	}, []point.Index{4}...)
	m.AddNaturalFrequency(2)
	m.AddNodeForce(2, force.NodeDim2{
		Fx: 10000.0,
		Fy: 10000.0,
	}, []point.Index{4}...)
	err := m.Solve()
	if err == nil {
		f7 := -26098.
		b, _, err := m.GetLocalForce(1, element.Index(7))
		if err != nil {
			return
		}
		if math.Abs((f7-b.Fx)/f7) > 0.01 {
			return
		} else {
			t.Errorf("axial force for beam 7 is %v cannot be equal without some data. Expected = %v", f7, b.Fx)
		}
	}
}
func TestGenTrussWothoutOneForFail2(t *testing.T) {
	var m model.Dim2
	m.AddPoint(point.Dim2{Index: 2, X: -0.8660254, Y: 0.})
	m.AddPoint(point.Dim2{Index: 1, X: 0., Y: 0.})
	m.AddPoint(point.Dim2{Index: 3, X: 0.8660254, Y: 0.})
	m.AddPoint(point.Dim2{Index: 4, X: 0., Y: -1.5})
	m.AddPoint(point.Dim2{Index: 40, X: 10., Y: 0.0})
	m.AddElement([]element.Elementer{
		element.NewBeam(7, 4, 2),
	}...)
	m.AddTrussProperty(7)
	m.AddTrussProperty(9, 8)
	m.AddSupport(support.FixedDim2(), 1)
	m.AddSupport(support.FixedDim2(), 3)
	m.AddSupport(support.FixedDim2(), 2)
	m.AddShape(shape.Shape{
		A: 300e-6,
	}, []element.Index{7, 9}...)
	m.AddShape(shape.Shape{
		A: 300e-6,
	}, []element.Index{8}...)
	m.AddMaterial(material.Linear{
		E:  2e11,
		Ro: 78500,
	}, []element.Index{9, 7}...)
	m.AddMaterial(material.Linear{
		E:  2e11,
		Ro: 78500,
	}, []element.Index{8}...)
	m.AddNodeForce(1, force.NodeDim2{
		Fy: -80000.0,
	}, []point.Index{4}...)
	m.AddNaturalFrequency(2)
	m.AddNodeForce(2, force.NodeDim2{
		Fx: 10000.0,
		Fy: 10000.0,
	}, []point.Index{4}...)
	err := m.Solve()
	if err == nil {
		f7 := -26098.
		b, _, err := m.GetLocalForce(1, element.Index(7))
		if err != nil {
			return
		}
		if math.Abs((f7-b.Fx)/f7) > 0.01 {
			return
		} else {
			t.Errorf("axial force for beam 7 is %v cannot be equal without some data. Expected = %v", f7, b.Fx)
		}
	}
}
func TestGenTrussWothoutOneForFail3(t *testing.T) {
	var m model.Dim2
	m.AddPoint(point.Dim2{Index: 2, X: -0.8660254, Y: 0.})
	m.AddPoint(point.Dim2{Index: 1, X: 0., Y: 0.})
	m.AddPoint(point.Dim2{Index: 3, X: 0.8660254, Y: 0.})
	m.AddPoint(point.Dim2{Index: 4, X: 0., Y: -1.5})
	m.AddPoint(point.Dim2{Index: 40, X: 10., Y: 0.0})
	m.AddElement([]element.Elementer{
		element.NewBeam(8, 4, 1),
		element.NewBeam(9, 4, 3),
	}...)
	m.AddTrussProperty(7)
	m.AddTrussProperty(9, 8)
	m.AddSupport(support.FixedDim2(), 1)
	m.AddSupport(support.FixedDim2(), 3)
	m.AddSupport(support.FixedDim2(), 2)
	m.AddShape(shape.Shape{
		A: 300e-6,
	}, []element.Index{7, 9}...)
	m.AddShape(shape.Shape{
		A: 300e-6,
	}, []element.Index{8}...)
	m.AddMaterial(material.Linear{
		E:  2e11,
		Ro: 78500,
	}, []element.Index{9, 7}...)
	m.AddMaterial(material.Linear{
		E:  2e11,
		Ro: 78500,
	}, []element.Index{8}...)
	m.AddNodeForce(1, force.NodeDim2{
		Fy: -80000.0,
	}, []point.Index{4}...)
	m.AddNaturalFrequency(2)
	m.AddNodeForce(2, force.NodeDim2{
		Fx: 10000.0,
		Fy: 10000.0,
	}, []point.Index{4}...)
	err := m.Solve()
	if err == nil {
		f7 := -26098.
		b, _, err := m.GetLocalForce(1, element.Index(7))
		if err != nil {
			return
		}
		if math.Abs((f7-b.Fx)/f7) > 0.01 {
			return
		} else {
			t.Errorf("axial force for beam 7 is %v cannot be equal without some data. Expected = %v", f7, b.Fx)
		}
	}
}
func TestGenTrussWothoutOneForFail4(t *testing.T) {
	var m model.Dim2
	m.AddPoint(point.Dim2{Index: 2, X: -0.8660254, Y: 0.})
	m.AddPoint(point.Dim2{Index: 1, X: 0., Y: 0.})
	m.AddPoint(point.Dim2{Index: 3, X: 0.8660254, Y: 0.})
	m.AddPoint(point.Dim2{Index: 4, X: 0., Y: -1.5})
	m.AddPoint(point.Dim2{Index: 40, X: 10., Y: 0.0})
	m.AddElement([]element.Elementer{
		element.NewBeam(8, 4, 1),
		element.NewBeam(9, 4, 3),
	}...)
	m.AddElement([]element.Elementer{
		element.NewBeam(7, 4, 2),
	}...)
	m.AddTrussProperty(9, 8)
	m.AddSupport(support.FixedDim2(), 1)
	m.AddSupport(support.FixedDim2(), 3)
	m.AddSupport(support.FixedDim2(), 2)
	m.AddShape(shape.Shape{
		A: 300e-6,
	}, []element.Index{7, 9}...)
	m.AddShape(shape.Shape{
		A: 300e-6,
	}, []element.Index{8}...)
	m.AddMaterial(material.Linear{
		E:  2e11,
		Ro: 78500,
	}, []element.Index{9, 7}...)
	m.AddMaterial(material.Linear{
		E:  2e11,
		Ro: 78500,
	}, []element.Index{8}...)
	m.AddNodeForce(1, force.NodeDim2{
		Fy: -80000.0,
	}, []point.Index{4}...)
	m.AddNaturalFrequency(2)
	m.AddNodeForce(2, force.NodeDim2{
		Fx: 10000.0,
		Fy: 10000.0,
	}, []point.Index{4}...)
	err := m.Solve()
	if err == nil {
		f7 := -26098.
		b, _, err := m.GetLocalForce(1, element.Index(7))
		if err != nil {
			return
		}
		if math.Abs((f7-b.Fx)/f7) > 0.01 {
			return
		} else {
			t.Errorf("axial force for beam 7 is %v cannot be equal without some data. Expected = %v", f7, b.Fx)
		}
	}
}
func TestGenTrussWothoutOneForFail5(t *testing.T) {
	var m model.Dim2
	m.AddPoint(point.Dim2{Index: 2, X: -0.8660254, Y: 0.})
	m.AddPoint(point.Dim2{Index: 1, X: 0., Y: 0.})
	m.AddPoint(point.Dim2{Index: 3, X: 0.8660254, Y: 0.})
	m.AddPoint(point.Dim2{Index: 4, X: 0., Y: -1.5})
	m.AddPoint(point.Dim2{Index: 40, X: 10., Y: 0.0})
	m.AddElement([]element.Elementer{
		element.NewBeam(8, 4, 1),
		element.NewBeam(9, 4, 3),
	}...)
	m.AddElement([]element.Elementer{
		element.NewBeam(7, 4, 2),
	}...)
	m.AddTrussProperty(7)
	m.AddSupport(support.FixedDim2(), 1)
	m.AddSupport(support.FixedDim2(), 3)
	m.AddSupport(support.FixedDim2(), 2)
	m.AddShape(shape.Shape{
		A: 300e-6,
	}, []element.Index{7, 9}...)
	m.AddShape(shape.Shape{
		A: 300e-6,
	}, []element.Index{8}...)
	m.AddMaterial(material.Linear{
		E:  2e11,
		Ro: 78500,
	}, []element.Index{9, 7}...)
	m.AddMaterial(material.Linear{
		E:  2e11,
		Ro: 78500,
	}, []element.Index{8}...)
	m.AddNodeForce(1, force.NodeDim2{
		Fy: -80000.0,
	}, []point.Index{4}...)
	m.AddNaturalFrequency(2)
	m.AddNodeForce(2, force.NodeDim2{
		Fx: 10000.0,
		Fy: 10000.0,
	}, []point.Index{4}...)
	err := m.Solve()
	if err == nil {
		f7 := -26098.
		b, _, err := m.GetLocalForce(1, element.Index(7))
		if err != nil {
			return
		}
		if math.Abs((f7-b.Fx)/f7) > 0.01 {
			return
		} else {
			t.Errorf("axial force for beam 7 is %v cannot be equal without some data. Expected = %v", f7, b.Fx)
		}
	}
}
func TestGenTrussWothoutOneForFail6(t *testing.T) {
	var m model.Dim2
	m.AddPoint(point.Dim2{Index: 2, X: -0.8660254, Y: 0.})
	m.AddPoint(point.Dim2{Index: 1, X: 0., Y: 0.})
	m.AddPoint(point.Dim2{Index: 3, X: 0.8660254, Y: 0.})
	m.AddPoint(point.Dim2{Index: 4, X: 0., Y: -1.5})
	m.AddPoint(point.Dim2{Index: 40, X: 10., Y: 0.0})
	m.AddElement([]element.Elementer{
		element.NewBeam(8, 4, 1),
		element.NewBeam(9, 4, 3),
	}...)
	m.AddElement([]element.Elementer{
		element.NewBeam(7, 4, 2),
	}...)
	m.AddTrussProperty(7)
	m.AddTrussProperty(9, 8)
	m.AddSupport(support.FixedDim2(), 3)
	m.AddSupport(support.FixedDim2(), 2)
	m.AddShape(shape.Shape{
		A: 300e-6,
	}, []element.Index{7, 9}...)
	m.AddShape(shape.Shape{
		A: 300e-6,
	}, []element.Index{8}...)
	m.AddMaterial(material.Linear{
		E:  2e11,
		Ro: 78500,
	}, []element.Index{9, 7}...)
	m.AddMaterial(material.Linear{
		E:  2e11,
		Ro: 78500,
	}, []element.Index{8}...)
	m.AddNodeForce(1, force.NodeDim2{
		Fy: -80000.0,
	}, []point.Index{4}...)
	m.AddNaturalFrequency(2)
	m.AddNodeForce(2, force.NodeDim2{
		Fx: 10000.0,
		Fy: 10000.0,
	}, []point.Index{4}...)
	err := m.Solve()
	if err == nil {
		f7 := -26098.
		b, _, err := m.GetLocalForce(1, element.Index(7))
		if err != nil {
			return
		}
		if math.Abs((f7-b.Fx)/f7) > 0.01 {
			return
		} else {
			t.Errorf("axial force for beam 7 is %v cannot be equal without some data. Expected = %v", f7, b.Fx)
		}
	}
}
func TestGenTrussWothoutOneForFail7(t *testing.T) {
	var m model.Dim2
	m.AddPoint(point.Dim2{Index: 2, X: -0.8660254, Y: 0.})
	m.AddPoint(point.Dim2{Index: 1, X: 0., Y: 0.})
	m.AddPoint(point.Dim2{Index: 3, X: 0.8660254, Y: 0.})
	m.AddPoint(point.Dim2{Index: 4, X: 0., Y: -1.5})
	m.AddPoint(point.Dim2{Index: 40, X: 10., Y: 0.0})
	m.AddElement([]element.Elementer{
		element.NewBeam(8, 4, 1),
		element.NewBeam(9, 4, 3),
	}...)
	m.AddElement([]element.Elementer{
		element.NewBeam(7, 4, 2),
	}...)
	m.AddTrussProperty(7)
	m.AddTrussProperty(9, 8)
	m.AddSupport(support.FixedDim2(), 1)
	m.AddShape(shape.Shape{
		A: 300e-6,
	}, []element.Index{7, 9}...)
	m.AddShape(shape.Shape{
		A: 300e-6,
	}, []element.Index{8}...)
	m.AddMaterial(material.Linear{
		E:  2e11,
		Ro: 78500,
	}, []element.Index{9, 7}...)
	m.AddMaterial(material.Linear{
		E:  2e11,
		Ro: 78500,
	}, []element.Index{8}...)
	m.AddNodeForce(1, force.NodeDim2{
		Fy: -80000.0,
	}, []point.Index{4}...)
	m.AddNaturalFrequency(2)
	m.AddNodeForce(2, force.NodeDim2{
		Fx: 10000.0,
		Fy: 10000.0,
	}, []point.Index{4}...)
	err := m.Solve()
	if err == nil {
		f7 := -26098.
		b, _, err := m.GetLocalForce(1, element.Index(7))
		if err != nil {
			return
		}
		if math.Abs((f7-b.Fx)/f7) > 0.01 {
			return
		} else {
			t.Errorf("axial force for beam 7 is %v cannot be equal without some data. Expected = %v", f7, b.Fx)
		}
	}
}
func TestGenTrussWothoutOneForFail8(t *testing.T) {
	var m model.Dim2
	m.AddPoint(point.Dim2{Index: 2, X: -0.8660254, Y: 0.})
	m.AddPoint(point.Dim2{Index: 1, X: 0., Y: 0.})
	m.AddPoint(point.Dim2{Index: 3, X: 0.8660254, Y: 0.})
	m.AddPoint(point.Dim2{Index: 4, X: 0., Y: -1.5})
	m.AddPoint(point.Dim2{Index: 40, X: 10., Y: 0.0})
	m.AddElement([]element.Elementer{
		element.NewBeam(8, 4, 1),
		element.NewBeam(9, 4, 3),
	}...)
	m.AddElement([]element.Elementer{
		element.NewBeam(7, 4, 2),
	}...)
	m.AddTrussProperty(7)
	m.AddTrussProperty(9, 8)
	m.AddSupport(support.FixedDim2(), 1)
	m.AddSupport(support.FixedDim2(), 3)
	m.AddSupport(support.FixedDim2(), 2)
	m.AddShape(shape.Shape{
		A: 300e-6,
	}, []element.Index{8}...)
	m.AddMaterial(material.Linear{
		E:  2e11,
		Ro: 78500,
	}, []element.Index{9, 7}...)
	m.AddMaterial(material.Linear{
		E:  2e11,
		Ro: 78500,
	}, []element.Index{8}...)
	m.AddNodeForce(1, force.NodeDim2{
		Fy: -80000.0,
	}, []point.Index{4}...)
	m.AddNaturalFrequency(2)
	m.AddNodeForce(2, force.NodeDim2{
		Fx: 10000.0,
		Fy: 10000.0,
	}, []point.Index{4}...)
	err := m.Solve()
	if err == nil {
		f7 := -26098.
		b, _, err := m.GetLocalForce(1, element.Index(7))
		if err != nil {
			return
		}
		if math.Abs((f7-b.Fx)/f7) > 0.01 {
			return
		} else {
			t.Errorf("axial force for beam 7 is %v cannot be equal without some data. Expected = %v", f7, b.Fx)
		}
	}
}
func TestGenTrussWothoutOneForFail9(t *testing.T) {
	var m model.Dim2
	m.AddPoint(point.Dim2{Index: 2, X: -0.8660254, Y: 0.})
	m.AddPoint(point.Dim2{Index: 1, X: 0., Y: 0.})
	m.AddPoint(point.Dim2{Index: 3, X: 0.8660254, Y: 0.})
	m.AddPoint(point.Dim2{Index: 4, X: 0., Y: -1.5})
	m.AddPoint(point.Dim2{Index: 40, X: 10., Y: 0.0})
	m.AddElement([]element.Elementer{
		element.NewBeam(8, 4, 1),
		element.NewBeam(9, 4, 3),
	}...)
	m.AddElement([]element.Elementer{
		element.NewBeam(7, 4, 2),
	}...)
	m.AddTrussProperty(7)
	m.AddTrussProperty(9, 8)
	m.AddSupport(support.FixedDim2(), 1)
	m.AddSupport(support.FixedDim2(), 3)
	m.AddSupport(support.FixedDim2(), 2)
	m.AddShape(shape.Shape{
		A: 300e-6,
	}, []element.Index{7, 9}...)
	m.AddMaterial(material.Linear{
		E:  2e11,
		Ro: 78500,
	}, []element.Index{9, 7}...)
	m.AddMaterial(material.Linear{
		E:  2e11,
		Ro: 78500,
	}, []element.Index{8}...)
	m.AddNodeForce(1, force.NodeDim2{
		Fy: -80000.0,
	}, []point.Index{4}...)
	m.AddNaturalFrequency(2)
	m.AddNodeForce(2, force.NodeDim2{
		Fx: 10000.0,
		Fy: 10000.0,
	}, []point.Index{4}...)
	err := m.Solve()
	if err == nil {
		f7 := -26098.
		b, _, err := m.GetLocalForce(1, element.Index(7))
		if err != nil {
			return
		}
		if math.Abs((f7-b.Fx)/f7) > 0.01 {
			return
		} else {
			t.Errorf("axial force for beam 7 is %v cannot be equal without some data. Expected = %v", f7, b.Fx)
		}
	}
}
func TestGenTrussWothoutOneForFail10(t *testing.T) {
	var m model.Dim2
	m.AddPoint(point.Dim2{Index: 2, X: -0.8660254, Y: 0.})
	m.AddPoint(point.Dim2{Index: 1, X: 0., Y: 0.})
	m.AddPoint(point.Dim2{Index: 3, X: 0.8660254, Y: 0.})
	m.AddPoint(point.Dim2{Index: 4, X: 0., Y: -1.5})
	m.AddPoint(point.Dim2{Index: 40, X: 10., Y: 0.0})
	m.AddElement([]element.Elementer{
		element.NewBeam(8, 4, 1),
		element.NewBeam(9, 4, 3),
	}...)
	m.AddElement([]element.Elementer{
		element.NewBeam(7, 4, 2),
	}...)
	m.AddTrussProperty(7)
	m.AddTrussProperty(9, 8)
	m.AddSupport(support.FixedDim2(), 1)
	m.AddSupport(support.FixedDim2(), 3)
	m.AddSupport(support.FixedDim2(), 2)
	m.AddShape(shape.Shape{
		A: 300e-6,
	}, []element.Index{7, 9}...)
	m.AddShape(shape.Shape{
		A: 300e-6,
	}, []element.Index{8}...)
	m.AddMaterial(material.Linear{
		E:  2e11,
		Ro: 78500,
	}, []element.Index{8}...)
	m.AddNodeForce(1, force.NodeDim2{
		Fy: -80000.0,
	}, []point.Index{4}...)
	m.AddNaturalFrequency(2)
	m.AddNodeForce(2, force.NodeDim2{
		Fx: 10000.0,
		Fy: 10000.0,
	}, []point.Index{4}...)
	err := m.Solve()
	if err == nil {
		f7 := -26098.
		b, _, err := m.GetLocalForce(1, element.Index(7))
		if err != nil {
			return
		}
		if math.Abs((f7-b.Fx)/f7) > 0.01 {
			return
		} else {
			t.Errorf("axial force for beam 7 is %v cannot be equal without some data. Expected = %v", f7, b.Fx)
		}
	}
}
func TestGenTrussWothoutOneForFail11(t *testing.T) {
	var m model.Dim2
	m.AddPoint(point.Dim2{Index: 2, X: -0.8660254, Y: 0.})
	m.AddPoint(point.Dim2{Index: 1, X: 0., Y: 0.})
	m.AddPoint(point.Dim2{Index: 3, X: 0.8660254, Y: 0.})
	m.AddPoint(point.Dim2{Index: 4, X: 0., Y: -1.5})
	m.AddPoint(point.Dim2{Index: 40, X: 10., Y: 0.0})
	m.AddElement([]element.Elementer{
		element.NewBeam(8, 4, 1),
		element.NewBeam(9, 4, 3),
	}...)
	m.AddElement([]element.Elementer{
		element.NewBeam(7, 4, 2),
	}...)
	m.AddTrussProperty(7)
	m.AddTrussProperty(9, 8)
	m.AddSupport(support.FixedDim2(), 1)
	m.AddSupport(support.FixedDim2(), 3)
	m.AddSupport(support.FixedDim2(), 2)
	m.AddShape(shape.Shape{
		A: 300e-6,
	}, []element.Index{7, 9}...)
	m.AddShape(shape.Shape{
		A: 300e-6,
	}, []element.Index{8}...)
	m.AddMaterial(material.Linear{
		E:  2e11,
		Ro: 78500,
	}, []element.Index{9, 7}...)
	m.AddNodeForce(1, force.NodeDim2{
		Fy: -80000.0,
	}, []point.Index{4}...)
	m.AddNaturalFrequency(2)
	m.AddNodeForce(2, force.NodeDim2{
		Fx: 10000.0,
		Fy: 10000.0,
	}, []point.Index{4}...)
	err := m.Solve()
	if err == nil {
		f7 := -26098.
		b, _, err := m.GetLocalForce(1, element.Index(7))
		if err != nil {
			return
		}
		if math.Abs((f7-b.Fx)/f7) > 0.01 {
			return
		} else {
			t.Errorf("axial force for beam 7 is %v cannot be equal without some data. Expected = %v", f7, b.Fx)
		}
	}
}
func TestGenTrussWothoutOneForFail12(t *testing.T) {
	var m model.Dim2
	m.AddPoint(point.Dim2{Index: 2, X: -0.8660254, Y: 0.})
	m.AddPoint(point.Dim2{Index: 1, X: 0., Y: 0.})
	m.AddPoint(point.Dim2{Index: 3, X: 0.8660254, Y: 0.})
	m.AddPoint(point.Dim2{Index: 4, X: 0., Y: -1.5})
	m.AddPoint(point.Dim2{Index: 40, X: 10., Y: 0.0})
	m.AddElement([]element.Elementer{
		element.NewBeam(8, 4, 1),
		element.NewBeam(9, 4, 3),
	}...)
	m.AddElement([]element.Elementer{
		element.NewBeam(7, 4, 2),
	}...)
	m.AddTrussProperty(7)
	m.AddTrussProperty(9, 8)
	m.AddSupport(support.FixedDim2(), 1)
	m.AddSupport(support.FixedDim2(), 3)
	m.AddSupport(support.FixedDim2(), 2)
	m.AddShape(shape.Shape{
		A: 300e-6,
	}, []element.Index{7, 9}...)
	m.AddShape(shape.Shape{
		A: 300e-6,
	}, []element.Index{8}...)
	m.AddMaterial(material.Linear{
		E:  2e11,
		Ro: 78500,
	}, []element.Index{9, 7}...)
	m.AddMaterial(material.Linear{
		E:  2e11,
		Ro: 78500,
	}, []element.Index{8}...)
	m.AddNaturalFrequency(2)
	m.AddNodeForce(2, force.NodeDim2{
		Fx: 10000.0,
		Fy: 10000.0,
	}, []point.Index{4}...)
	err := m.Solve()
	if err == nil {
		f7 := -26098.
		b, _, err := m.GetLocalForce(1, element.Index(7))
		if err != nil {
			return
		}
		if math.Abs((f7-b.Fx)/f7) > 0.01 {
			return
		} else {
			t.Errorf("axial force for beam 7 is %v cannot be equal without some data. Expected = %v", f7, b.Fx)
		}
	}
}
