package finiteElement

import (
	"github.com/Konstantin8105/GoFea/dof"
	"github.com/Konstantin8105/GoFea/material"
	"github.com/Konstantin8105/GoFea/point"
	"github.com/Konstantin8105/GoFea/shape"
	"github.com/Konstantin8105/GoLinAlg/linAlg"
)

// TrussDim2 - truss on 2D interpratation
type TrussDim2 struct {
	Material material.Linear
	Shape    shape.Shape
	Points   [2]point.Dim2
}

// GetCoordinateTransformation - matrix of transform between local and global system coordinate
func (f *TrussDim2) GetCoordinateTransformation(tr *linAlg.Matrix64) {
	lenght := point.LenghtDim2(f.Points)
	lambdaXX := (f.Points[1].X - f.Points[0].X) / lenght
	lambdaXY := (f.Points[1].Y - f.Points[0].Y) / lenght

	tr.SetNewSize(2, 4)
	tr.Set(0, 0, lambdaXX)
	tr.Set(0, 1, lambdaXY)
	tr.Set(1, 2, lambdaXX)
	tr.Set(1, 3, lambdaXY)
}

// GetStiffinerK - matrix of stiffiner
func (f *TrussDim2) GetStiffinerK(kr *linAlg.Matrix64) {
	lenght := point.LenghtDim2(f.Points)
	EFL := f.Material.E * f.Shape.A / lenght

	kr.SetNewSize(2, 2)
	kr.Set(0, 0, EFL)
	kr.Set(1, 0, -EFL)
	kr.Set(0, 1, -EFL)
	kr.Set(1, 1, EFL)
}

// GetMassMr - matrix mass of finite element
func (f *TrussDim2) GetMassMr(mr *linAlg.Matrix64) {
	mu := f.Shape.A * f.Material.Ro
	lenght := point.LenghtDim2(f.Points)
	mul3 := lenght / 3.0 * mu
	mul6 := lenght / 6.0 * mu

	mr.SetNewSize(2, 2)
	mr.Set(0, 0, mul3)
	mr.Set(1, 0, mul6)
	mr.Set(0, 1, mul6)
	mr.Set(1, 1, mul3)
}

// GetDoF - return numbers for degree of freedom
func (f *TrussDim2) GetDoF(degrees *dof.DoF) (axes []dof.AxeNumber) {
	var Axe [2][]dof.AxeNumber
	Axe[0] = degrees.GetDoF(f.Points[0].Index)
	Axe[1] = degrees.GetDoF(f.Points[1].Index)

	inx := 0
	axes = make([]dof.AxeNumber, 4, 4)
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			axes[inx] = Axe[i][j]
			inx++
		}
	}
	return
}

// GetStiffinerGlobalK - global matrix of stiffiner
func GetStiffinerGlobalK(f FiniteElementer, degree *dof.DoF, info Information) (linAlg.Matrix64, []dof.AxeNumber) {
	klocal := linAlg.NewMatrix64bySize(4, 4)
	f.GetStiffinerK(&klocal)

	Tr := linAlg.NewMatrix64bySize(4, 4)
	f.GetCoordinateTransformation(&Tr)

	kor := klocal.MultiplyTtKT(Tr)

	axes := f.GetDoF(degree)

	var removePosition []int
	for i := 0; i < len(axes); i++ {
		found := false
		for j := 0; j < len(axes); j++ {
			if kor.Get(i, j) != 0.0 {
				found = true
				break
			}
		}
		if found {
			continue
		}
		removePosition = append(removePosition, i)
	}

	if info == WithoutZeroStiffiner {
		// remove row and column from global stiffiner
		kor.RemoveRowAndColumn(removePosition...)
		// remove column from axes
		dof.RemoveIndexes(&axes, removePosition...)
	}

	return kor, axes
}

// GetStiffinerGlobalMass - global matrix of mass
func GetStiffinerGlobalMass(f FiniteElementer, degree *dof.DoF, info Information) (linAlg.Matrix64, []dof.AxeNumber) {
	mlocal := linAlg.NewMatrix64bySize(4, 4)
	f.GetMassMr(&mlocal)

	Tr := linAlg.NewMatrix64bySize(4, 4)
	f.GetCoordinateTransformation(&Tr)

	mor := mlocal.MultiplyTtKT(Tr)

	axes := f.GetDoF(degree)

	var removePosition []int
	for i := 0; i < len(axes); i++ {
		found := false
		for j := 0; j < len(axes); j++ {
			if mor.Get(i, j) != 0.0 {
				found = true
				break
			}
		}
		if found {
			continue
		}
		removePosition = append(removePosition, i)
	}

	if info == WithoutZeroStiffiner {
		// remove row and column from global stiffiner
		mor.RemoveRowAndColumn(removePosition...)
		// remove column from axes
		dof.RemoveIndexes(&axes, removePosition...)
	}

	return mor, axes
}
