package finiteElement

import (
	"github.com/Konstantin8105/GoFea/dof"
	"github.com/Konstantin8105/GoFea/material"
	"github.com/Konstantin8105/GoFea/point"
	"github.com/Konstantin8105/GoFea/shape"
	"github.com/Konstantin8105/GoLinAlg/matrix"
)

// TrussDim2 - truss on 2D interpratation
type TrussDim2 struct {
	Material material.Linear
	Shape    shape.Shape
	Points   [2]point.Dim2
}

// GetCoordinateTransformation - matrix of transform between local and global system coordinate
func (f *TrussDim2) GetCoordinateTransformation(tr *matrix.T64) {
	lenght := point.LenghtDim2(f.Points)
	lambdaXX := (f.Points[1].X - f.Points[0].X) / lenght
	lambdaXY := (f.Points[1].Y - f.Points[0].Y) / lenght
	/*
		tr.SetNewSize(2, 4)
		tr.Set(0, 0, lambdaXX)
		tr.Set(0, 1, lambdaXY)
		tr.Set(1, 2, lambdaXX)
		tr.Set(1, 3, lambdaXY)
	*/
	tr.SetNewSize(6, 6)
	tr.Set(0, 0, lambdaXX)
	tr.Set(0, 1, lambdaXY)
	tr.Set(1, 0, -lambdaXY)
	tr.Set(1, 1, lambdaXX)
	tr.Set(2, 2, 1.0)

	tr.Set(3, 3, lambdaXX)
	tr.Set(3, 4, lambdaXY)
	tr.Set(4, 3, -lambdaXY)
	tr.Set(4, 4, lambdaXX)
	tr.Set(5, 5, 1.0)

}

// GetStiffinerK - matrix of stiffiner
func (f *TrussDim2) GetStiffinerK(kr *matrix.T64) {
	lenght := point.LenghtDim2(f.Points)
	EFL := f.Material.E * f.Shape.A / lenght
	/*
		kr.SetNewSize(2, 2)
		kr.Set(0, 0, EFL)
		kr.Set(1, 0, -EFL)
		kr.Set(0, 1, -EFL)
		kr.Set(1, 1, EFL)
	*/
	kr.SetNewSize(6, 6)
	kr.Set(0, 0, EFL)
	kr.Set(0, 3, -EFL)
	kr.Set(3, 0, -EFL)
	kr.Set(3, 3, EFL)
}

// GetMassMr - matrix mass of finite element
func (f *TrussDim2) GetMassMr(mr *matrix.T64) {
	mu := f.Shape.A * f.Material.Ro
	lenght := point.LenghtDim2(f.Points)
	mul3 := lenght / 3.0 * mu
	mul6 := lenght / 6.0 * mu
	/*
		mr.SetNewSize(2, 2)
		mr.Set(0, 0, mul3)
		mr.Set(1, 0, mul6)
		mr.Set(0, 1, mul6)
		mr.Set(1, 1, mul3)
	*/
	mr.SetNewSize(6, 6)
	mr.Set(0, 0, mul3)
	mr.Set(0, 3, mul6)
	mr.Set(3, 0, mul6)
	mr.Set(3, 3, mul3)
}

// GetPotentialGr - matrix potential loads for linear buckling
func (f *TrussDim2) GetPotentialGr(gr *matrix.T64, localAxialForce float64) {
	lenght := point.LenghtDim2(f.Points)
	/*
		gr.SetNewSize(2, 2)
		gr.Set(0, 0, 1.0/lenght)
		gr.Set(1, 1, 1.0/lenght)
	*/

	NL := localAxialForce / lenght
	// TODO check somewhere lenght cannot by zero

	gr.SetNewSize(6, 6)
	gr.Set(1, 1, NL)
	gr.Set(1, 4, -NL)
	gr.Set(4, 1, -NL)
	gr.Set(4, 4, NL)
}

// GetDoF - return numbers for degree of freedom
func (f *TrussDim2) GetDoF(degrees *dof.DoF) (axes []dof.AxeNumber) {
	var Axe [2][]dof.AxeNumber
	Axe[0] = degrees.GetDoF(f.Points[0].Index)
	Axe[1] = degrees.GetDoF(f.Points[1].Index)
	/*
		inx := 0
		axes = make([]dof.AxeNumber, 4, 4)
		for i := 0; i < 2; i++ {
			for j := 0; j < 2; j++ {
				axes[inx] = Axe[i][j]
				inx++
			}
		}
	*/
	inx := 0
	axes = make([]dof.AxeNumber, 6, 6)
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			axes[inx] = Axe[i][j]
			inx++
		}
	}
	return
}

// GetStiffinerGlobalK - global matrix of stiffiner
func GetStiffinerGlobalK(f FiniteElementer, degree *dof.DoF, info Information) (matrix.T64, []dof.AxeNumber) {
	klocal := matrix.NewMatrix64bySize(4, 4)
	f.GetStiffinerK(&klocal)

	Tr := matrix.NewMatrix64bySize(4, 4)
	f.GetCoordinateTransformation(&Tr)

	kor := klocal.MultiplyTtKT(Tr)

	axes := f.GetDoF(degree)

	if info == WithoutZeroStiffiner {
		RemoveZeros(&kor, &axes)
	}

	return kor, axes
}

// GetGlobalMass - global matrix of mass
func GetGlobalMass(f FiniteElementer, degree *dof.DoF, info Information) (matrix.T64, []dof.AxeNumber) {
	mlocal := matrix.NewMatrix64bySize(4, 4)
	f.GetMassMr(&mlocal)

	Tr := matrix.NewMatrix64bySize(4, 4)
	f.GetCoordinateTransformation(&Tr)

	mor := mlocal.MultiplyTtKT(Tr)

	axes := f.GetDoF(degree)

	if info == WithoutZeroStiffiner {
		RemoveZeros(&mor, &axes)
	}

	return mor, axes
}

/*
// GetGlobalPotential - global matrix of mass
func GetGlobalPotential(f FiniteElementer, degree *dof.DoF, info Information) (linAlg.Matrix64, []dof.AxeNumber) {
	mlocal := linAlg.NewMatrix64bySize(4, 4)
	f.GetPotentialGr(&mlocal, -1)
	panic("TODO")

	Tr := linAlg.NewMatrix64bySize(4, 4)
	f.GetCoordinateTransformation(&Tr)

	mor := mlocal.MultiplyTtKT(Tr)

	axes := f.GetDoF(degree)

	if info == WithoutZeroStiffiner {
		RemoveZeros(&mor, &axes)
	}

	return mor, axes
}
*/

// RemoveZeros - remove columns, rows of matrix and columns of dof
func RemoveZeros(matrix *matrix.T64, axes *[]dof.AxeNumber) {
	var removePosition []int
	// TODO: len --> to matrix lenght
	// TODO: at the first check diagonal element
	for i := 0; i < len(*axes); i++ {
		found := false
		for j := 0; j < len(*axes); j++ {
			if (*matrix).Get(i, j) != 0.0 {
				found = true
				break
			}
		}
		if found {
			continue
		}
		removePosition = append(removePosition, i)
	}

	// TODO: can parallel
	// remove row and column from global stiffiner
	(*matrix).RemoveRowAndColumn(removePosition...)
	// remove column from axes
	dof.RemoveIndexes(axes, removePosition...)
}
