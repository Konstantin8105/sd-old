package model

import (
	"testing"

	"github.com/Konstantin8105/GoFea/input/element"
	"github.com/Konstantin8105/GoFea/input/material"
	"github.com/Konstantin8105/GoFea/input/point"
	"github.com/Konstantin8105/GoFea/input/shape"
)

type fake struct {
	index        element.Index
	pointIndexes []point.Index
}

func newFake(i element.Index, p0, p1 point.Index) (f fake) {
	f.index = i
	f.pointIndexes = append(f.pointIndexes, p0, p1)
	return
}
func (f fake) GetIndex() element.Index {
	return f.index
}
func (f fake) GetPointIndex() []point.Index {
	return f.pointIndexes
}
func (f fake) GetAmountPoint() int {
	return 2
}

func TestPanicFE1(t *testing.T) {
	var m Dim2
	m.AddPoint([]point.Dim2{
		{Index: 1},
		{Index: 2, X: 1.0, Y: 1.0},
		{Index: 3, X: 1.0, Y: 0.0},
	}...)
	m.AddShape(shape.Shape{A: 1.0}, []element.Index{1, 2}...)
	m.AddMaterial(material.Linear{E: 2e11, Ro: 78500}, []element.Index{1, 2}...)
	m.AddElement([]element.Elementer{
		newFake(1, 1, 2),
		newFake(2, 2, 3),
	}...)
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	_, _ = m.getFiniteElement(2)
}

func TestErrorCoordinate(t *testing.T) {
	var m Dim2
	m.AddPoint([]point.Dim2{
		{Index: 1},
		//	{Index: 2, X: 1.0, Y: 1.0},
		{Index: 3, X: 1.0, Y: 0.0},
	}...)
	m.AddShape(shape.Shape{A: 1.0}, []element.Index{1, 2}...)
	m.AddMaterial(material.Linear{E: 2e11, Ro: 78500}, []element.Index{1, 2}...)
	m.AddElement([]element.Elementer{
		element.NewBeam(1, 1, 2),
		element.NewBeam(2, 2, 3),
	}...)
	_, err := m.getFiniteElement(2)
	if err == nil {
		t.Errorf("Cannot fount - wrong coordinate")
	}
}

func TestErrorElements(t *testing.T) {
	var m Dim2
	m.AddPoint([]point.Dim2{
		{Index: 1},
		{Index: 2, X: 1.0, Y: 1.0},
		{Index: 3, X: 1.0, Y: 0.0},
	}...)
	m.AddShape(shape.Shape{A: 1.0}, []element.Index{1, 2}...)
	m.AddMaterial(material.Linear{E: 2e11, Ro: 78500}, []element.Index{1, 2}...)
	m.AddElement([]element.Elementer{
		element.NewBeam(1, 1, 2),
		//element.NewBeam(2, 2, 3),
	}...)
	_, err := m.getFiniteElement(2)
	if err == nil {
		t.Errorf("Cannot fount - wrong element")
	}
}
