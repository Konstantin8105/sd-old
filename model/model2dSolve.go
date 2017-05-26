package model

import (
	"fmt"

	"github.com/Konstantin8105/GoFea/dof"
	"github.com/Konstantin8105/GoFea/element"
	"github.com/Konstantin8105/GoFea/finiteElement"
	"github.com/Konstantin8105/GoFea/utils"
	"github.com/Konstantin8105/GoLinAlg/linAlg"
	"github.com/Konstantin8105/GoLinAlg/linAlg/solver"
)

// Solve - solving finite element
func (m *Dim2) Solve() (err error) {

	for caseNumber := 0; caseNumber < len(m.forceCases); caseNumber++ {

		// TODO : check everything
		// TODO : sort  everything
		// TODO : compress loads by number

		// Generate degree of freedom in global system
		var degreeGlobal []dof.AxeNumber
		dofSystem := dof.NewBeam(m.beams, dof.Dim2d)
		for _, beam := range m.beams {
			_, degreeLocal := m.globalStiffinerForBeam(beam.Index, &dofSystem, finiteElement.WithoutZeroStiffiner)
			degreeGlobal = append(degreeGlobal, degreeLocal...)
		}
		{
			is := dof.ConvertToInt(degreeGlobal)
			utils.UniqueInt(&is)
			degreeGlobal = dof.ConvertToAxe(is)
		}

		// Generate global stiffiner matrix
		stiffinerKGlobal := linAlg.NewMatrix64bySize(len(degreeGlobal), len(degreeGlobal))
		mapIndex := dof.NewMapIndex(&degreeGlobal)
		for _, beam := range m.beams {
			klocal, degreeLocal := m.globalStiffinerForBeam(beam.Index, &dofSystem, finiteElement.WithoutZeroStiffiner)
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

		// create load vector
		loads := linAlg.NewMatrix64bySize(len(degreeGlobal), 1)
		for _, node := range m.forceCases[caseNumber].nodeForces {
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
		//fmt.Println("degreeGlobal = ", degreeGlobal)
		//fmt.Printf("K global = \n%s\n", stiffinerKGlobal)
		//fmt.Printf("Load vector = \n%s\n", loads)

		// Solving system of linear equations for finding
		// the displacement in points in global system
		lu := solver.NewLUsolver(stiffinerKGlobal)
		x := lu.Solve(loads)

		fmt.Printf("Global displacement = \n%s\n", x)
		fmt.Println("degreeGlobal = ", degreeGlobal)
		for _, beam := range m.beams {
			klocal, degreeLocal := m.globalStiffinerForBeam(beam.Index, &dofSystem, finiteElement.FullInformation)
			fmt.Println("=============")
			fmt.Println("klocalGlobal = ", klocal)
			fmt.Println("degreeLocal = ", degreeLocal)
			globalDisplacement := make([]float64, len(degreeLocal))
			for i := 0; i < len(globalDisplacement); i++ {
				found := false
				for j := 0; j < len(degreeGlobal); j++ {
					if degreeLocal[i] == degreeGlobal[j] {
						found = true
						globalDisplacement[i] = x.Get(j, 0)
						break
					}
				}
				if !found {
					panic("Cannot found dof - MAY BE PINNED. Check")
				}
			}
			fmt.Println("globalDisplacement = ", globalDisplacement)

			inx := beam.Index
			material, err := m.getMaterial(inx)
			if err != nil {
				panic(fmt.Errorf("Cannot found material for beam #%v. Error = %v", inx, err))
			}
			shape, err := m.getShape(inx)
			if err != nil {
				panic(fmt.Errorf("Cannot found shape for beam #%v. Error = %v", inx, err))
			}
			coord, err := m.getCoordinate(inx)
			if err != nil {
				panic(fmt.Errorf("Cannot calculate lenght for beam #%v. Error = %v", inx, err))
			}
			tr := finiteElement.TrussDim2{
				Material: material,
				Shape:    shape,
				Points:   coord,
			}
			t := linAlg.NewMatrix64bySize(10, 10)
			tr.GetCoordinateTransformation(&t)
			fmt.Println("tr.glo --", t)

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

			kk := linAlg.NewMatrix64bySize(10, 10)
			tr.GetStiffinerK(&kk)
			fmt.Println("klocalll -->", kk)

			var localForce []float64
			for i := 0; i < kk.GetRowSize(); i++ {
				sum := 0.0
				for j := 0; j < kk.GetRowSize(); j++ {
					sum += kk.Get(i, j) * localDisplacement[j]
				}
				localForce = append(localForce, sum)
			}

			fmt.Println("localForce = ", localForce)

		}
	}

	return nil
}

func (m *Dim2) globalStiffinerForBeam(inx element.BeamIndex, dofSystem *dof.DoF, info finiteElement.Information) (localK linAlg.Matrix64, degrees []dof.AxeNumber) {

	material, err := m.getMaterial(inx)
	if err != nil {
		panic(fmt.Errorf("Cannot found material for beam #%v. Error = %v", inx, err))
	}
	shape, err := m.getShape(inx)
	if err != nil {
		panic(fmt.Errorf("Cannot found shape for beam #%v. Error = %v", inx, err))
	}
	coord, err := m.getCoordinate(inx)
	if err != nil {
		panic(fmt.Errorf("Cannot calculate lenght for beam #%v. Error = %v", inx, err))
	}
	if m.isTruss(inx) {
		tr := finiteElement.TrussDim2{
			Material: material,
			Shape:    shape,
			Points:   coord,
		}
		return tr.GetStiffinerGlobalK(dofSystem, info)
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
	return localK, degrees
}
