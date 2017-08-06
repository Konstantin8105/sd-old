package finiteElement_test

import (
	"testing"

	"github.com/Konstantin8105/GoFea/input/material"
	"github.com/Konstantin8105/GoFea/input/point"
	"github.com/Konstantin8105/GoFea/input/shape"
	"github.com/Konstantin8105/GoFea/solver/finiteElement"
	"github.com/Konstantin8105/GoLinAlg/matrix"
)

func TestBeam2dSymmetrical(t *testing.T) {
	f := finiteElement.BeamDim2{
		Material: material.Linear{
			E: 1.0,
		},
		Shape: shape.Shape{
			A:   2.0,
			Izz: 3.0,
		},
		Points: [2]point.Dim2{
			{Index: 1, X: 0.000, Y: 0.000},
			{Index: 2, X: 0.000, Y: 4.000},
		},
	}
	functions := []func(*matrix.T64){
		f.GetStiffinerK,
		f.GetMassMr,
		func(m *matrix.T64) {
			f.GetPotentialGr(m, 5.000)
		},
	}
	for i := range functions {
		m := matrix.NewMatrix64bySize(1, 1)
		(functions[i])(&m)
		if m.GetRowSize() != 6 {
			t.Errorf("Size is not 6 elements")
		}
		if m.GetColumnSize() != 6 {
			t.Errorf("Size is not 6 elements")
		}

		for i := 0; i < 6; i++ {
			for j := 0; j < 6; j++ {
				if m.Get(i, j) != m.Get(j, i) {
					t.Errorf("Not symmetrical for [%v,%v]", i, j)
				}
			}
		}
	}
}
