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

/*
func TestLinearBucklingBeams(t *testing.T) {
	var m model.Dim2

	m.AddPoint(point.Dim2{Index: 1, X: 0.000, Y: 0.400})
	m.AddPoint(point.Dim2{Index: 2, X: 0.200, Y: 0.400})
	m.AddPoint(point.Dim2{Index: 3, X: 0.400, Y: 0.400})
	m.AddPoint(point.Dim2{Index: 4, X: 0.400, Y: 0.200})
	m.AddPoint(point.Dim2{Index: 5, X: 0.400, Y: 0.000})
	m.AddPoint(point.Dim2{Index: 6, X: 0.600, Y: 0.400})
	m.AddPoint(point.Dim2{Index: 7, X: 0.800, Y: 0.400})
	m.AddPoint(point.Dim2{Index: 8, X: 0.800, Y: 0.200})
	m.AddPoint(point.Dim2{Index: 9, X: 0.800, Y: 0.000})

	m.AddElement(element.NewBeam(1, 1, 2))
	m.AddElement(element.NewBeam(2, 2, 3))
	m.AddElement(element.NewBeam(3, 3, 4))
	m.AddElement(element.NewBeam(4, 4, 5))
	m.AddElement(element.NewBeam(5, 3, 6))
	m.AddElement(element.NewBeam(6, 6, 7))
	m.AddElement(element.NewBeam(7, 7, 8)) //Pin at the begin
	m.AddElement(element.NewBeam(8, 8, 9))

	m.AddSupport(support.Dim2{Dx: support.Fix, Dy: support.Fix}, 1)
	m.AddSupport(support.FixedDim2(), 5)
	m.AddSupport(support.FixedDim2(), 9)

	m.AddShape(shape.Shape{
		A:   24e-4,
		Izz: 32e-8,
	}, []element.Index{1, 2, 3, 4, 7, 8}...)
	m.AddShape(shape.Shape{
		A:   48e-4,
		Izz: 64e-8,
	}, []element.Index{5, 6}...)

	m.AddMaterial(material.Linear{
		E:  2e11,
		Ro: 78500,
	}, []element.Index{1, 2, 3, 4, 5, 6, 7, 8}...)

	m.AddNodeForce(1, force.NodeDim2{
		Fy: -0.8,
	}, 3)
	m.AddNodeForce(1, force.NodeDim2{
		Fy: -1.0,
	}, 7)

	m.AddLinearBuckling(1)

	err := m.Solve()
	if err != nil {
		t.Errorf("Cannot solving. error = %v", err)
	}

	fmt.Println("What is ALPHA in input data")
	fmt.Println("Add pin in element")
	fmt.Println("Critical buckling loads must be 8860.4660 N")
	fmt.Println("Add checking buckling shape")

}
*/
/*
func TestLinearBucklingGFrame(t *testing.T) {
	var m model.Dim2

	m.AddPoint(point.Dim2{Index: 1, X: 0.000, Y: 0.000})
	m.AddPoint(point.Dim2{Index: 2, X: 0.000, Y: 1.000})
	m.AddPoint(point.Dim2{Index: 3, X: 0.000, Y: 2.000})
	m.AddPoint(point.Dim2{Index: 4, X: 0.000, Y: 3.000})
	m.AddPoint(point.Dim2{Index: 5, X: 0.000, Y: 4.000})
	m.AddPoint(point.Dim2{Index: 6, X: 3.000, Y: 4.000})

	m.AddElement(element.NewBeam(1, 1, 2))
	m.AddElement(element.NewBeam(2, 2, 3))
	m.AddElement(element.NewBeam(3, 3, 4))
	m.AddElement(element.NewBeam(4, 4, 5))
	m.AddElement(element.NewBeam(5, 5, 6))

	m.AddSupport(support.FixedDim2(), 1)
	m.AddSupport(support.Dim2{Dx: support.Fix, Dy: support.Fix}, 6)

	J := 1.0e-8
	m.AddShape(shape.Shape{
		A:   24e-4,
		Izz: J,
	}, []element.Index{1, 2, 3, 4}...)
	m.AddShape(shape.Shape{
		A:   24e-4,
		Izz: 2.0 * J,
	}, []element.Index{5}...)

	E := 2.0e11
	m.AddMaterial(material.Linear{
		E:  E,
		Ro: 78500,
	}, []element.Index{1, 2, 3, 4, 5}...)

	m.AddNodeForce(1, force.NodeDim2{
		Fy: -1.0,
	}, 5)

	m.AddLinearBuckling(1)

	err := m.Solve()
	if err != nil {
		t.Errorf("Cannot solving. error = %v", err)
	}

	fmt.Println("Critical force from = ", 2.0*E*J, " to ", 2.04*E*J, " N")
}
*/
/*
func TestTframe(t *testing.T) {
	var m model.Dim2

	m.AddPoint(point.Dim2{Index: 1, X: 0.000, Y: 0.400})
	m.AddPoint(point.Dim2{Index: 2, X: 0.200, Y: 0.400})
	m.AddPoint(point.Dim2{Index: 3, X: 0.400, Y: 0.000})
	m.AddPoint(point.Dim2{Index: 4, X: 0.400, Y: 0.200})
	m.AddPoint(point.Dim2{Index: 5, X: 0.400, Y: 0.400})
	m.AddPoint(point.Dim2{Index: 6, X: 0.700, Y: 0.400})
	m.AddPoint(point.Dim2{Index: 7, X: 1.000, Y: 0.400})

	m.AddElement(element.NewBeam(1, 1, 2))
	m.AddElement(element.NewBeam(2, 2, 5))
	m.AddElement(element.NewBeam(3, 3, 4))
	m.AddElement(element.NewBeam(4, 4, 5))
	m.AddElement(element.NewBeam(5, 5, 6))
	m.AddElement(element.NewBeam(6, 6, 7))

	m.AddSupport(support.FixedDim2(), []point.Index{1, 3, 7}...)

	m.AddShape(shape.Shape{
		A:   24e-4,
		Izz: 72e-8,
	}, []element.Index{1, 2, 3, 4, 5, 6}...)

	m.AddMaterial(material.Linear{
		E:  2e11,
		Ro: 78500,
	}, []element.Index{1, 2, 3, 4, 5, 6}...)

	m.AddNodeForce(1, force.NodeDim2{
		//Fx: 1000.0,
		Fy: 1000.0,
	}, []point.Index{2, 6}...)

	m.AddNaturalFrequency(1)

	err := m.Solve()
	if err != nil {
		t.Errorf("Cannot solving. error = %v", err)
	}

	hz, _ := m.GetNaturalFrequency(1)
	fmt.Println("Hz:", hz)
	fmt.Println("Expeted : 26.777 sec-1")
}
*/
/*
func TestTwoBeams(t *testing.T) {
	var m model.Dim2

	L := 0.100
	m.AddPoint(point.Dim2{Index: 1, X: 0.000, Y: 0.000})
	m.AddPoint(point.Dim2{Index: 2, X: 1.0 * L, Y: 0.000})
	m.AddPoint(point.Dim2{Index: 3, X: 2.0 * L, Y: 0.000})

	m.AddElement(element.NewBeam(1, 1, 2))
	m.AddElement(element.NewBeam(2, 2, 3))

	m.AddSupport(support.FixedDim2(), []point.Index{1, 3}...)

	J := 40e-8
	A := 30e-4
	m.AddShape(shape.Shape{
		A:   A,
		Izz: J,
	}, []element.Index{1, 2}...)

	E := 2e11
	Ro := 78500.
	m.AddMaterial(material.Linear{
		E:  E,
		Ro: Ro,
	}, []element.Index{1, 2}...)

	m.AddNaturalFrequency(1)

	err := m.Solve()
	if err != nil {
		t.Errorf("Cannot solving. error = %v", err)
	}

	hz, _ := m.GetNaturalFrequency(1)
	fmt.Println("Hz:", hz)
	w1 := 4.90 / L / L * math.Sqrt(E*J/(Ro*A)) // 5.68
	//w2 := 20.49 / L / L * math.Sqrt(E*J/(Ro*A))
	fmt.Println("Expeted : w1 = ", w1, " sec-1")
	//fmt.Println("Expeted : w2 = ", w2, " sec-1")
	f1 := w1 / 2.0 / math.Pi
	//f2 := w2 / 2.0 / math.Pi
	fmt.Println("Expeted : f1 = ", f1, " Hz")
	//fmt.Println("Expeted : f2 = ", f2, " Hz")

}
*/
func TestBendBeam(t *testing.T) {
	var m model.Dim2

	m.AddPoint(point.Dim2{Index: 1, X: 0.000, Y: 0.000})
	m.AddPoint(point.Dim2{Index: 2, X: 1.000, Y: 0.000})
	m.AddPoint(point.Dim2{Index: 3, X: 2.000, Y: 0.000})
	m.AddPoint(point.Dim2{Index: 4, X: 3.000, Y: 0.000})

	m.AddElement(element.NewBeam(1, 1, 2))
	m.AddElement(element.NewBeam(2, 2, 3))
	m.AddElement(element.NewBeam(3, 3, 4))

	m.AddSupport(support.FixedDim2(), 4)

	m.AddShape(shape.Shape{
		A:   20e-4,
		Izz: 40e-6,
	}, []element.Index{1, 2, 3}...)

	m.AddMaterial(material.Linear{
		E:  2e11,
		Ro: 78500,
	}, []element.Index{1, 2, 3}...)

	m.AddNodeForce(1, force.NodeDim2{
		M: 3000.0,
	}, []point.Index{2}...)

	m.AddNodeForce(1, force.NodeDim2{
		Fy: -4000.0,
	}, []point.Index{1}...)

	m.AddNodeForce(1, force.NodeDim2{
		Fy: 7000.0,
	}, []point.Index{3}...)

	err := m.Solve()
	if err != nil {
		t.Errorf("Cannot solving. error = %v", err)
	}

	r, err := m.GetReaction(1, 4)
	if err != nil {
		t.Errorf("Cannot get reaction")
	}

	fy := -3000.0
	if math.Abs((r.Fy-fy)/fy) > 1e-5 {
		t.Errorf("Not correct reaction = %v. Expected = %v", r.Fy, fy)
	}
}

func TestOsscilatorAxial(t *testing.T) {
	var m model.Dim2

	E := 2e11
	F := 20e-4
	Q := 2500.0
	L := 1.500

	m.AddPoint(point.Dim2{Index: 1, X: 0.000, Y: 0.000})
	m.AddPoint(point.Dim2{Index: 2, X: L, Y: 0.000})

	m.AddElement(element.NewBeam(1, 1, 2))

	m.AddSupport(support.FixedDim2(), 1)

	m.AddShape(shape.Shape{
		A:   F,
		Izz: 40e-6,
	}, []element.Index{1}...)

	m.AddMaterial(material.Linear{
		E:  E,
		Ro: 78500,
	}, []element.Index{1}...)

	m.AddNodeForce(1, force.NodeDim2{
		Fx: Q,
	}, []point.Index{2}...)

	m.AddNaturalFrequency(1)

	err := m.Solve()
	if err != nil {
		t.Errorf("Cannot solving. error = %v", err)
	}

	w := math.Sqrt(9.80 * E * F / (Q * L))
	f := w / (2.0 * math.Pi)
	{
		hz1 := f
		actualHz, err := m.GetNaturalFrequency(1)
		if err != nil {
			t.Errorf("Cannot found natural frequency for case 1. Error = %v", err)
		}
		{
			var found bool
			for i := range actualHz {
				if math.Abs((hz1-actualHz[i])/hz1) < 0.015 {
					found = true
				}
			}
			if !found {
				t.Errorf("Natural frequency calculated not correct = %vHz. Expected = %vHz", actualHz, hz1)
			}
		}
	}

}

func TestOsscilatorBend(t *testing.T) {
	var m model.Dim2

	E := 2e11
	F := 20e-4
	Q := 2500.0
	L := 1.500
	J := 60e-6

	m.AddPoint(point.Dim2{Index: 1, X: 0.000, Y: 0.000})
	m.AddPoint(point.Dim2{Index: 2, X: L * 1. / 5., Y: 0.000})
	m.AddPoint(point.Dim2{Index: 3, X: L * 2. / 5., Y: 0.000})
	m.AddPoint(point.Dim2{Index: 4, X: L * 3. / 5., Y: 0.000})
	m.AddPoint(point.Dim2{Index: 5, X: L * 4. / 5., Y: 0.000})
	m.AddPoint(point.Dim2{Index: 6, X: L, Y: 0.000})

	m.AddElement(element.NewBeam(1, 1, 2))
	m.AddElement(element.NewBeam(2, 2, 3))
	m.AddElement(element.NewBeam(3, 3, 4))
	m.AddElement(element.NewBeam(4, 4, 5))
	m.AddElement(element.NewBeam(5, 5, 6))

	m.AddSupport(support.FixedDim2(), 1)

	m.AddShape(shape.Shape{
		A:   F,
		Izz: J,
	}, []element.Index{1, 2, 3, 4, 5}...)

	m.AddMaterial(material.Linear{
		E:  E,
		Ro: 78500,
	}, []element.Index{1, 2, 3, 4, 5}...)

	m.AddNodeForce(1, force.NodeDim2{
		Fy: Q,
	}, []point.Index{6}...)

	m.AddNaturalFrequency(1)

	err := m.Solve()
	if err != nil {
		t.Errorf("Cannot solving. error = %v", err)
	}

	w := math.Sqrt(3.0 * 9.80 * E * J / (Q * math.Pow(L, 3.0)))
	f := w / (2.0 * math.Pi)
	{
		hz1 := f
		actualHz, err := m.GetNaturalFrequency(1)
		if err != nil {
			t.Errorf("Cannot found natural frequency for case 1. Error = %v", err)
		}
		{
			var found bool
			for i := range actualHz {
				if math.Abs((hz1-actualHz[i])/hz1) < 0.015 {
					found = true
				}
			}
			if !found {
				t.Errorf("Natural frequency calculated not correct = %vHz. Expected = %vHz", actualHz, hz1)
			}
		}
	}

}
