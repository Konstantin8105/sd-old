package model

import (
	"fmt"

	"github.com/Konstantin8105/GoFea/dof"
	"github.com/Konstantin8105/GoFea/element"
	"github.com/Konstantin8105/GoFea/finiteElement"
	"github.com/Konstantin8105/GoFea/utils"
	"github.com/Konstantin8105/GoLinAlg/matrix"
	"github.com/Konstantin8105/GoLinAlg/solver"
)

func (m *Dim2) solveCase(forceCase *forceCase2d) error {

	fmt.Println("case = ", forceCase.indexCase)
	// TODO : check everything
	// TODO : sort  everything
	// TODO : compress loads by number

	// Generate degree of freedom in global system
	var degreeGlobal []dof.AxeNumber
	dofSystem := dof.NewBeam(m.elements, dof.Dim2d)
	for _, ele := range m.elements {
		switch ele.(type) {
		case element.Beam:
			beam := ele.(element.Beam)
			fe := m.getBeamFiniteElement(beam.Index)
			_, degreeLocal := finiteElement.GetStiffinerGlobalK(fe, &dofSystem, finiteElement.WithoutZeroStiffiner)
			degreeGlobal = append(degreeGlobal, degreeLocal...)
		default:
			panic("")
		}
	}
	{
		is := dof.ConvertToInt(degreeGlobal)
		utils.UniqueInt(&is)
		degreeGlobal = dof.ConvertToAxe(is)
	}

	// Create convertor index to axe
	mapIndex := dof.NewMapIndex(&degreeGlobal)

	// Generate global stiffiner matrix [Ko]
	stiffinerKGlobal := m.convertFromLocalToGlobalSystem(&degreeGlobal, &dofSystem, &mapIndex, finiteElement.GetStiffinerGlobalK)

	// Create load vector
	loads := matrix.NewMatrix64bySize(len(degreeGlobal), 1)
	for _, node := range forceCase.nodeForces {
		for _, inx := range node.pointIndexes {
			d := dofSystem.GetDoF(inx)
			if node.nodeForce.Fx != 0.0 {
				h, err := mapIndex.GetByAxe(d[0])
				if err == nil {
					loads.Set(h, 0, node.nodeForce.Fx)
				}
			}
			if node.nodeForce.Fy != 0.0 {
				h, err := mapIndex.GetByAxe(d[1])
				if err == nil {
					loads.Set(h, 0, node.nodeForce.Fy)
				}
			}
			if node.nodeForce.M != 0.0 {
				h, err := mapIndex.GetByAxe(d[2])
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
			d := dofSystem.GetDoF(inx)
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
				g, err := mapIndex.GetByAxe(result[i])
				if err != nil {
					continue
				}
				for j := 0; j < len(degreeGlobal); j++ {
					h, err := mapIndex.GetByAxe(degreeGlobal[j])
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
	fmt.Println("degreeGlobal = ", degreeGlobal)
	fmt.Printf("K global = \n%s\n", stiffinerKGlobal)
	fmt.Printf("Load vector = \n%s\n", loads)

	// Solving system of linear equations for finding
	// the displacement in points in global system
	// TODO: if you have nonlinear elements, then we can use
	// TODO: one global stiffiner matrix for all cases
	lu := solver.NewLUsolver(stiffinerKGlobal)
	x := lu.Solve(loads)
	// TODO: rename global vector of displacement

	fmt.Printf("Global displacement = \n%s\n", x)
	//fmt.Println("degreeGlobal = ", degreeGlobal)
	for _, ele := range m.elements {
		switch ele.(type) {
		case element.Beam:
			beam := ele.(element.Beam)
			fe := m.getBeamFiniteElement(beam.Index)
			/*klocal,*/ _, degreeLocal := finiteElement.GetStiffinerGlobalK(fe, &dofSystem, finiteElement.FullInformation)
			//fmt.Println("=============")
			//fmt.Println("klocalGlobal = ", klocal)
			//fmt.Println("degreeLocal = ", degreeLocal)
			globalDisplacement := make([]float64, len(degreeLocal))
			// if not found in global displacement, then it is a pinned
			// in local stiffiner matrix - than row and column is zero
			// for avoid collisian - we put a zero
			for i := 0; i < len(globalDisplacement); i++ {
				for j := 0; j < len(degreeGlobal); j++ {
					if degreeLocal[i] == degreeGlobal[j] {
						globalDisplacement[i] = x.Get(j, 0)
						break
					}
				}
			}
			//fmt.Println("globalDisplacement = ", globalDisplacement)

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
