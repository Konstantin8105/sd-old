package model

import (
	"fmt"
	"math"

	"github.com/Konstantin8105/GoFea/solver/finiteElement"
	"github.com/Konstantin8105/GoLinAlg/matrix"
	"github.com/Konstantin8105/GoLinAlg/solver"
)

func (m *Dim2) solveNaturalFrequency(forceCase *forceCase2d) error {

	lu, err := m.getLUStiffinerKGlobal()
	if err != nil {
		return err
	}

	// Generate global mass matrix [Mo]
	//n := stiffinerKGlobal.GetRowSize()
	n := len(m.degreeInGlobalMatrix)
	massGlobal, err := m.convertFromLocalToGlobalSystem(&m.degreeInGlobalMatrix, &m.degreeForPoint, &m.indexsInGlobalMatrix, finiteElement.GetGlobalMass)
	if err != nil {
		return err
	}
	// m.convertFromLocalToGlobalSystem(&degreeGlobal, &dofSystem, &mapIndex, finiteElement.GetGlobalMass)
	//  linAlg.NewMatrix64bySize(n, n)

	// TODO: Add to matrix mass the nodal mass
	for _, p := range forceCase.nodeForces {
		//	for _, inx := range node.pointIndexes {
		d := m.degreeForPoint.GetDoF(p.pointIndex)
		if p.nodeForce.Fx != 0.0 {
			h, err := m.indexsInGlobalMatrix.GetByAxe(d[0])
			if err == nil {
				massGlobal.Set(h, h, massGlobal.Get(h, h)+math.Abs(p.nodeForce.Fx))
			}
		}
		if p.nodeForce.Fy != 0.0 {
			h, err := m.indexsInGlobalMatrix.GetByAxe(d[1])
			if err == nil {
				massGlobal.Set(h, h, massGlobal.Get(h, h)+math.Abs(p.nodeForce.Fy))
			}
		}
		// TODO: Moment haven`t mass ???
		// TODO: Check
		//if node.nodeForce.M != 0.0 {
		//	h, err := mapIndex.GetByAxe(d[2])
		//	if err == nil {
		//		massGlobal.Set(h, h, massGlobal.Get(h, h)+math.Abs(node.nodeForce.M))
		//		fmt.Println("Add M to mass")
		//	}
		//}
		//	}
	}

	//TODO: CHECKUING GRAVITY TO MATRIX MASS
	for i := 0; i < massGlobal.GetRowSize(); i++ {
		for j := 0; j < massGlobal.GetColumnSize(); j++ {
			massGlobal.Set(i, j, massGlobal.Get(i, j)/9.806)
		}
	}
	// TODO: ADD to mass WITH OR WITOUT SELFWEIGHT

	// Calculate matrix [H] = [Ko]^-1 * [Mo]
	//if stiffinerKGlobal.GetRowSize() != stiffinerKGlobal.GetColumnSize() {
	//	panic("Not correct size of global stiffiner matrix")
	//}
	//fmt.Println("GlobalMass = ", massGlobal)
	Ho := matrix.NewMatrix64bySize(n, n)
	buffer := matrix.NewMatrix64bySize(n, 1)
	for i := 0; i < n; i++ {
		// Create vertical vector from [Mo]
		for j := 0; j < n; j++ {
			buffer.Set(j, 0, massGlobal.Get(j, i))
		}
		// Calculation
		result := lu.Solve(buffer)
		// Add vector to [Ho]
		for j := 0; j < n; j++ {
			Ho.Set(j, i, result.Get(j, 0))
		}
	}
	//fmt.Println("[Ho] = ", Ho)
	{
		// TODO: check
		// Remove zero rows and columns
		var removePosition []int
		// TODO: len --> to matrix length
		// TODO: at the first check diagonal element
		for i := 0; i < Ho.GetRowSize(); i++ {
			found := false
			for j := 0; j < Ho.GetRowSize(); j++ {
				if Ho.Get(i, j) != 0.0 {
					found = true
					break
				}
			}
			if found {
				continue
			}
			removePosition = append(removePosition, i)
		}
		Ho.RemoveRowAndColumn(removePosition...)
	}

	// Calculation of natural frequency
	eigen := solver.NewEigen(Ho)
	//fmt.Println("lambda       = ", eigen.GetRealEigenvalues())
	//fmt.Println("eigenvectors = ", eigen.GetV())
	//fmt.Println("getD = ", eigen.GetD())

	// TODO: fix for avoid strange frequency some is too small or too big
	value := eigen.GetRealEigenvalues()
	for _, v := range value {
		freq := math.Sqrt(1.0/v) / 2.0 / math.Pi
		fmt.Printf("f = %.5v Hz\n", freq)
		_ = freq
	}
	// TODO: need add modal mass values for natural frquency calculation
	return nil
}
