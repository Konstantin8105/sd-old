package model

import (
	"fmt"

	"github.com/Konstantin8105/GoFea/input/element"
	"github.com/Konstantin8105/GoFea/output/displacement"
	"github.com/Konstantin8105/GoFea/solver/dof"
	"github.com/Konstantin8105/GoFea/solver/finiteElement"
	"github.com/Konstantin8105/GoLinAlg/matrix"
	"github.com/Konstantin8105/GoLinAlg/solver"
)

func (m *Dim2) solveCase(forceCase *forceCase2d) error {

	fmt.Println("case = ", forceCase.indexCase)
	// TODO : check everything
	// TODO : sort  everything
	// TODO : compress loads by number

	// Generate global stiffiner matrix [Ko]
	stiffinerKGlobal := m.convertFromLocalToGlobalSystem(&m.degreeInGlobalMatrix, &m.degreeForPoint, &m.indexsInGlobalMatrix, finiteElement.GetStiffinerGlobalK)

	// Create load vector
	loads := matrix.NewMatrix64bySize(len(m.degreeInGlobalMatrix), 1)
	for _, node := range forceCase.nodeForces {
		for _, inx := range node.pointIndexes {
			d := m.degreeForPoint.GetDoF(inx)
			if node.nodeForce.Fx != 0.0 {
				h, err := m.indexsInGlobalMatrix.GetByAxe(d[0])
				if err == nil {
					loads.Set(h, 0, node.nodeForce.Fx)
				}
			}
			if node.nodeForce.Fy != 0.0 {
				h, err := m.indexsInGlobalMatrix.GetByAxe(d[1])
				if err == nil {
					loads.Set(h, 0, node.nodeForce.Fy)
				}
			}
			if node.nodeForce.M != 0.0 {
				h, err := m.indexsInGlobalMatrix.GetByAxe(d[2])
				if err == nil {
					loads.Set(h, 0, node.nodeForce.M)
				}
			}
		}
	}

	// Create array degree for support
	// and modify the global stiffiner matrix
	// and load vector
	for _, sup := range m.supports {
		for _, inx := range sup.pointIndexes {
			d := m.degreeForPoint.GetDoF(inx)
			var result []dof.AxeNumber
			if sup.support.Dx == true {
				result = append(result, d[0])
			}
			if sup.support.Dy == true {
				result = append(result, d[1])
			}
			if sup.support.M == true {
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
					stiffinerKGlobal.Set(g, h, 0.0)
					stiffinerKGlobal.Set(h, g, 0.0)
				}
				stiffinerKGlobal.Set(g, g, 1.0)
				// modify load vector on support
				loads.Set(g, 0, 0.0)
			}
		}
	}
	fmt.Println("degreeGlobal = ", m.degreeInGlobalMatrix)
	fmt.Printf("K global = \n%s\n", stiffinerKGlobal)
	fmt.Printf("Load vector = \n%s\n", loads)

	// Solving system of linear equations for finding
	// the displacement in points in global system
	// TODO: if you have nonlinear elements, then we can use
	// TODO: one global stiffiner matrix for all cases
	lu := solver.NewLUsolver(stiffinerKGlobal)
	globalDisp := lu.Solve(loads)
	// TODO: rename global vector of displacement

	fmt.Printf("Global displacement = \n%s\n", globalDisp)

	// global displacement for points
	for _, p := range m.points {
		axes := m.degreeForPoint.GetDoF(p.Index)
		var disp displacement.Dim2
		disp.Index = p.Index
		for i := range axes {
			for j := 0; j < len(m.degreeInGlobalMatrix); j++ {
				// TODO : only for 2d
				if axes[i] == m.degreeInGlobalMatrix[j] {
					if i == 0 {
						disp.Dx = globalDisp.Get(j, 0)
					}
					if i == 1 {
						disp.Dy = globalDisp.Get(j, 0)
					}
					if i == 2 {
						disp.Dm = globalDisp.Get(j, 0)
					}
				}
			}
		}
		forceCase.globalDisplacements = append(forceCase.globalDisplacements, disp)
		fmt.Println("disp -- > ", disp)
	}
	//fmt.Println("degreeGlobal = ", degreeGlobal)
	for _, ele := range m.elements {
		switch ele.(type) {
		case element.Beam:
			beam := ele.(element.Beam)
			fe := m.getBeamFiniteElement(beam.Index)
			/*klocal,*/ _, degreeLocal := finiteElement.GetStiffinerGlobalK(fe, &m.degreeForPoint, finiteElement.FullInformation)
			//fmt.Println("=============")
			//fmt.Println("klocalGlobal = ", klocal)
			//fmt.Println("degreeLocal = ", degreeLocal)
			globalDisplacement := make([]float64, len(degreeLocal))
			// if not found in global displacement, then it is a pinned
			// in local stiffiner matrix - than row and column is zero
			// for avoid collisian - we put a zero
			for i := 0; i < len(globalDisplacement); i++ {
				for j := 0; j < len(m.degreeInGlobalMatrix); j++ {
					if degreeLocal[i] == m.degreeInGlobalMatrix[j] {
						globalDisplacement[i] = globalDisp.Get(j, 0)
						break
					}
				}
			}
			fmt.Println("globalDisplacement = ", globalDisplacement)

			t := matrix.NewMatrix64bySize(10, 10)
			fe.GetCoordinateTransformation(&t)
			//fmt.Println("tr.glo --", t)

			// Zo = T_t * Z
			var localDisplacement []float64
			for i := 0; i < t.GetRowSize(); i++ {
				sum := 0.0
				for j := 0; j < t.GetColumnSize(); j++ {
					sum += t.Get(i, j) * globalDisplacement[j]
				}
				localDisplacement = append(localDisplacement, sum)
			}
			fmt.Println("localDisplacement = ", localDisplacement)

			kk := matrix.NewMatrix64bySize(10, 10)
			fe.GetStiffinerK(&kk)
			//fmt.Println("klocalll -->", kk)

			var localForce []float64
			for i := 0; i < kk.GetRowSize(); i++ {
				sum := 0.0
				for j := 0; j < kk.GetRowSize(); j++ {
					sum += kk.Get(i, j) * localDisplacement[j]
				}
				localForce = append(localForce, sum)
			}
			fmt.Println("localForce = ", localForce)
			_ = localForce
		default:
			panic("")
		}
	}

	return nil
}
