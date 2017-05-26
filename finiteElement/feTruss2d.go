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

// GetCoordinateTransformation - record into buffer a matrix of transform from local to global system coordinate
func (f *TrussDim2) GetCoordinateTransformation(buffer *linAlg.Matrix64) {
	buffer.SetNewSize(2, 4)

	lenght := point.LenghtDim2(f.Points)

	lambdaXX := (f.Points[1].X - f.Points[0].X) / lenght
	lambdaXY := (f.Points[1].Y - f.Points[0].Y) / lenght

	buffer.Set(0, 0, lambdaXX)
	buffer.Set(0, 1, lambdaXY)
	buffer.Set(1, 2, lambdaXX)
	buffer.Set(1, 3, lambdaXY)
}

// GetStiffinerK - matrix of stiffiner
func (f *TrussDim2) GetStiffinerK(buffer *linAlg.Matrix64) {
	buffer.SetNewSize(2, 2)

	lenght := point.LenghtDim2(f.Points)

	EFL := f.Material.E * f.Shape.A / lenght

	buffer.Set(0, 0, EFL)
	buffer.Set(1, 0, -EFL)
	buffer.Set(0, 1, -EFL)
	buffer.Set(1, 1, EFL)
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

// GetStiffinerGlobalK - global matrix of siffiner
func (f *TrussDim2) GetStiffinerGlobalK(degree *dof.DoF, info Information) (linAlg.Matrix64, []dof.AxeNumber) {
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
