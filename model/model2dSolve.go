package model

import (
	"fmt"

	"github.com/Konstantin8105/GoFea/dof"
	"github.com/Konstantin8105/GoFea/finiteElement"
	"github.com/Konstantin8105/GoFea/linearAlgebra"
	"github.com/Konstantin8105/GoFea/utils"
)

// Solve - solving finite element
func (m *Dim2) Solve() (err error) {

	for caseNumber := 0; caseNumber < len(m.forceCases); caseNumber++ {

		// TODO : check everything
		// TODO : sort  everything
		// TODO : compress loads by number

		type stepSolving int
		const (
			prepareDegreeOfFreedom stepSolving = iota
			createGlobalStiffinerK
		)
		steps := []stepSolving{prepareDegreeOfFreedom, createGlobalStiffinerK}

		var degreeGlobal []dof.AxeNumber
		var stiffinerKGlobal linearAlgebra.Matrix
		var mapIndex dof.MapIndex

		var loads linearAlgebra.Vector

		for _, step := range steps {
			dofSystem := dof.NewBeam(m.beams, dof.Dim2d)

			if step == createGlobalStiffinerK {
				stiffinerKGlobal = linearAlgebra.NewSquareMatrix(len(degreeGlobal))
				mapIndex = dof.NewMapIndex(&degreeGlobal)
				loads = linearAlgebra.NewVector(len(degreeGlobal))
			}

			var klocal linearAlgebra.Matrix
			var degreeLocal []dof.AxeNumber
			for _, beam := range m.beams {
				material, err := m.getMaterial(beam.Index)
				if err != nil {
					return fmt.Errorf("Cannot found material for beam #%v. Error = %v", beam.Index, err)
				}
				shape, err := m.getShape(beam.Index)
				if err != nil {
					return fmt.Errorf("Cannot found shape for beam #%v. Error = %v", beam.Index, err)
				}
				coord, err := m.getCoordinate(beam.Index)
				if err != nil {
					return fmt.Errorf("Cannot calculate lenght for beam #%v. Error = %v", beam.Index, err)
				}
				if m.isTruss(beam.Index) {
					tr := finiteElement.TrussDim2{
						Material: material,
						Shape:    shape,
						Points:   coord,
					}
					klocal, degreeLocal = tr.GetStiffinerGlobalK(&dofSystem)
				} /* else {
					fe := finiteElement.BeamDim2{
						Material: material,
						Shape:    shape,
						Points:   coord,
					}
					err = fe.GetStiffinerK(&buffer)
					if err != nil {
						return err
					}
				}*/
				if step == prepareDegreeOfFreedom {
					degreeGlobal = append(degreeGlobal, degreeLocal...)
				}
				if step == createGlobalStiffinerK {
					// Add local stiffiner matrix to global matrix
					for i := 0; i < len(degreeLocal); i++ {
						g, err := mapIndex.GetByAxe(degreeLocal[i])
						if err != nil {
							continue
						}
						for j := 0; j < len(degreeLocal); j++ {
							h, err := mapIndex.GetByAxe(degreeLocal[j])
							if err != nil {
								continue
							}
							stiffinerKGlobal.Set(g, h, stiffinerKGlobal.Get(g, h)+klocal.Get(i, j))
						}
					}
				}
			}
			if step == prepareDegreeOfFreedom {
				is := dof.ConvertToInt(degreeGlobal)
				utils.UniqueInt(&is)
				degreeGlobal = dof.ConvertToAxe(is)
			}
			if step == createGlobalStiffinerK {
				// create load vector
				fmt.Println("Load case : ", m.forceCases[caseNumber].indexCase)
				for _, node := range m.forceCases[caseNumber].nodeForces {
					for _, inx := range node.pointIndexes {
						d := dofSystem.GetDoF(inx)
						if node.nodeForce.Fx != 0.0 {
							h, err := mapIndex.GetByAxe(d[0])
							if err == nil {
								loads.Set(h, node.nodeForce.Fx)
							}
						}
						if node.nodeForce.Fy != 0.0 {
							h, err := mapIndex.GetByAxe(d[1])
							if err == nil {
								loads.Set(h, node.nodeForce.Fy)
							}
						}
						if node.nodeForce.M != 0.0 {
							h, err := mapIndex.GetByAxe(d[2])
							if err == nil {
								loads.Set(h, node.nodeForce.M)
							}
						}
					}
				}

				// Create array degree for support
				var supportDegree []dof.AxeNumber
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
						supportDegree = append(supportDegree, result...)
					}
				}
				// supportDegree - created array degree for support

				// create global stiffiner matrix
				for i := 0; i < len(supportDegree); i++ {
					g, err := mapIndex.GetByAxe(supportDegree[i])
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
					loads.Set(g, 0.0)
				}
			}
		}
		fmt.Println("degreeGlobal = ", degreeGlobal)
		fmt.Printf("K global = \n%s\n", stiffinerKGlobal)
		fmt.Printf("Load vector = \n%s\n", loads)
	}

	return nil
}
