package finiteElement

import (
	"fmt"

	"github.com/Konstantin8105/GoFea/solver/dof"
	"github.com/Konstantin8105/GoLinAlg/matrix"
)

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
	if matrix.GetRowSize() != len(*axes) || matrix.GetColumnSize() != len(*axes) {
		panic(fmt.Errorf("Not correct input data.\nMatrix = %v\nAxes = %v", *matrix, *axes))
	}
	size := len(*axes)
	var removePosition []int
	for i := 0; i < size; i++ {
		found := false
		if (*matrix).Get(i, i) == 0.0 {
			for j := 0; j < size; j++ {
				if (*matrix).Get(i, j) != 0.0 {
					found = true
					goto end
				}
			}
		} else {
			found = true
		}
	end:
		if found {
			continue
		}
		removePosition = append(removePosition, i)
	}

	// remove row and column from global stiffiner
	(*matrix).RemoveRowAndColumn(removePosition...)
	// remove column from axes
	dof.RemoveIndexes(axes, removePosition...)
}
