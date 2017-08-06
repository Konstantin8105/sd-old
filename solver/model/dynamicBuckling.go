package model

import (
	"github.com/Konstantin8105/GoFea/solver/dof"
	"github.com/Konstantin8105/GoFea/solver/finiteElement"
	"github.com/Konstantin8105/GoLinAlg/matrix"
	"github.com/Konstantin8105/GoLinAlg/solver"
)

func (m *Dim2) solveLinearBuckling(forceCase *forceCase2d) error {

	_ = m.solveCase(forceCase)

	lu, err := m.getLUStiffinerKGlobal()
	if err != nil {
		return err
	}

	//potentialGlobal := m.convertFromLocalToGlobalSystem(&degreeGlobal, &dofSystem, &mapIndex, finiteElement.GetGlobalPotential)
	//potentialGlobal := matrix.NewMatrix64bySize(stiffinerKGlobal.GetRowSize(), stiffinerKGlobal.GetColumnSize())

	potentialGlobal := matrix.NewMatrix64bySize(len(m.degreeInGlobalMatrix), len(m.degreeInGlobalMatrix))

	n := len(m.degreeInGlobalMatrix)

	for _, ele := range m.elements {
		fe, err := m.getFiniteElement(ele.GetIndex())
		if err != nil {
			return err
		}

		//klocal,
		//_, degreeLocal := finiteElement.GetStiffinerGlobalK(fe, &dofSystem, finiteElement.FullInformation)
		_, degreeLocal := finiteElement.GetStiffinerGlobalK(fe, &m.degreeForPoint, finiteElement.WithoutZeroStiffiner)

		/*
				globalDisplacement := make([]float64, len(degreeLocal))
				// if not found in global displacement, then it is a pinned
				// in local stiffiner matrix - than row and column is zero
				// for avoid collisian - we put a zero
				for i := 0; i < len(globalDisplacement); i++ {
					for j := 0; j < len(m.degreeInGlobalMatrix); j++ {
						if degreeLocal[i] == m.degreeInGlobalMatrix[j] {
							globalDisplacement[i] = x.Get(j, 0)
							break
						}
					}
				}

				t := matrix.NewMatrix64bySize(10, 10)
				fe.GetCoordinateTransformation(&t)

				// Zo = T_t * Z
				var localDisplacement []float64
				for i := 0; i < t.GetRowSize(); i++ {
					sum := 0.0
					for j := 0; j < t.GetColumnSize(); j++ {
						sum += t.Get(i, j) * globalDisplacement[j]
					}
					localDisplacement = append(localDisplacement, sum)
				}

				kk := matrix.NewMatrix64bySize(10, 10)
				fe.GetStiffinerK(&kk)

				var localForce []float64
				for i := 0; i < kk.GetRowSize(); i++ {
					sum := 0.0
					for j := 0; j < kk.GetRowSize(); j++ {
						sum += kk.Get(i, j) * localDisplacement[j]
					}
					localForce = append(localForce, sum)
				}

			//fmt.Println("local Force = ", localForce)
			if localForce[0] > 0.0 && localForce[3] < 0.0 {
				// TODO : it is not correct , because uniform load can change
				//fmt.Println("Compress")
			} else {
				// TODO: testing
				localForce[0] = 0.0
			}
		*/

		b, e, err := forceCase.GetLocalForce(ele.GetIndex())
		if err != nil {
			return err
		}
		var axialForce float64
		if b.Fx > 0.0 && e.Fx < 0.0 {
			// TODO : it is not correct , because uniform load can change
			//fmt.Println("Compress")
			axialForce = b.Fx
		} /*else {
			// TODO: testing
			localForce[0] = 0.0
		}*/

		grLocal := matrix.NewMatrix64bySize(6, 6)
		fe.GetPotentialGr(&grLocal, axialForce) //localForce[0])

		// Add local stiffiner matrix to global matrix
		for i := 0; i < len(degreeLocal); i++ {
			g, err := m.indexsInGlobalMatrix.GetByAxe(degreeLocal[i])
			if err != nil {
				continue
			}
			for j := 0; j < len(degreeLocal); j++ {
				h, err := m.indexsInGlobalMatrix.GetByAxe(degreeLocal[j])
				if err != nil {
					continue
				}
				potentialGlobal.Set(g, h, potentialGlobal.Get(g, h)+grLocal.Get(i, j))
			}
		}
	}
	//fmt.Println("PotentialGlobal = ", potentialGlobal)

	// Create array degree for support
	// and modify the global stiffiner matrix
	// and load vector
	for _, s := range m.supports {
		d := m.degreeForPoint.GetDoF(s.pointIndex)
		var result []dof.AxeNumber
		if s.support.Dx {
			result = append(result, d[0])
		}
		if s.support.Dy {
			result = append(result, d[1])
		}
		if s.support.M {
			result = append(result, d[2])
		}
		// modify stiffiner matrix for correct
		// adding support
		for i := 0; i < len(result); i++ {
			g, err := m.indexsInGlobalMatrix.GetByAxe(result[i])
			if err != nil {
				continue
			}
			for j := 0; j < len(m.degreeInGlobalMatrix); j++ {
				h, err := m.indexsInGlobalMatrix.GetByAxe(m.degreeInGlobalMatrix[j])
				if err != nil {
					continue
				}
				potentialGlobal.Set(g, h, 0.0)
				potentialGlobal.Set(h, g, 0.0)
			}
			//potentialGlobal.Set(g, g, 1.0)
		}
	}

	//fmt.Println("PotentialGlobal = ", potentialGlobal)
	HoPotential := matrix.NewMatrix64bySize(n, n)
	bufferPotential := matrix.NewMatrix64bySize(n, 1)
	//fmt.Printf("lu = %#v\n", lu)
	for i := 0; i < n; i++ {
		// Create vertical vector from [Mo]
		for j := 0; j < n; j++ {
			bufferPotential.Set(j, 0, potentialGlobal.Get(j, i))
		}
		// Calculation
		result, err := lu.Solve(bufferPotential)
		if err != nil {
			return err
		}
		// Add vector to [Ho]
		for j := 0; j < n; j++ {
			HoPotential.Set(j, i, result.Get(j, 0))
		}
	}
	//fmt.Println("[HoPotential] = ", HoPotential)
	//fmt.Println("row1", HoPotential.GetRowSize())
	//fmt.Println("col1", HoPotential.GetColumnSize())
	{
		// TODO: check
		// Remove zero rows and columns
		var removePosition []int
		// TODO: len --> to matrix length
		// TODO: at the first check diagonal element
		for i := 0; i < HoPotential.GetRowSize(); i++ {
			found := false
			for j := 0; j < HoPotential.GetRowSize(); j++ {
				if HoPotential.Get(i, j) != 0.0 {
					found = true
					break
				}
			}
			if found {
				continue
			}
			removePosition = append(removePosition, i)
		}
		HoPotential.RemoveRowAndColumn(removePosition...)
	}
	// Calculation of

	// TODO add for tension beam - panic

	//fmt.Println("row2", HoPotential.GetRowSize())
	//fmt.Println("col2", HoPotential.GetColumnSize())
	eigenPotential := solver.NewEigen(HoPotential)
	//fmt.Println("lambda       = ", eigenPotential.GetRealEigenvalues())
	//fmt.Println("lambda Re    = ", eigenPotential.GetImagEigenvalues())
	//fmt.Println("eigenvectors = ", eigenPotential.GetV())
	//fmt.Println("getD = ", eigenPotential.GetD())

	// TODO: Remove strange results
	valueP := eigenPotential.GetRealEigenvalues()
	// fmt.Println("Linear buckling loads:")
	for _, v := range valueP {
		// fmt.Printf("v = %.5v\t\tP = %.5v\n", v, 1.0/v)
		// TODO sorting by absolute value
		forceCase.dynamicValue = append(forceCase.dynamicValue, 1.0/v)
	}

	return nil
}
