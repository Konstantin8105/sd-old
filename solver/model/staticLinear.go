package model

import (
	"fmt"

	"github.com/Konstantin8105/GoFea/input/element"
	"github.com/Konstantin8105/GoFea/output/displacement"
	"github.com/Konstantin8105/GoFea/output/forceLocal"
	"github.com/Konstantin8105/GoFea/output/reaction"
	"github.com/Konstantin8105/GoFea/solver/dof"
	"github.com/Konstantin8105/GoFea/solver/finiteElement"
	"github.com/Konstantin8105/GoLinAlg/matrix"
	"github.com/Konstantin8105/GoLinAlg/solver"
)

func (m *Dim2) solveCase(forceCase *forceCase2d) error {

	// Generate global stiffiner matrix [Ko]
	stiffinerKGlobal := m.convertFromLocalToGlobalSystem(&m.degreeInGlobalMatrix, &m.degreeForPoint, &m.indexsInGlobalMatrix, finiteElement.GetStiffinerGlobalK)

	// Create load vector
	loads := matrix.NewMatrix64bySize(len(m.degreeInGlobalMatrix), 1)
	for _, p := range forceCase.nodeForces {
		d := m.degreeForPoint.GetDoF(p.pointIndex)
		if p.nodeForce.Fx != 0.0 {
			h, err := m.indexsInGlobalMatrix.GetByAxe(d[0])
			if err == nil {
				loads.Set(h, 0, p.nodeForce.Fx)
			}
		}
		if p.nodeForce.Fy != 0.0 {
			h, err := m.indexsInGlobalMatrix.GetByAxe(d[1])
			if err == nil {
				loads.Set(h, 0, p.nodeForce.Fy)
			}
		}
		if p.nodeForce.M != 0.0 {
			h, err := m.indexsInGlobalMatrix.GetByAxe(d[2])
			if err == nil {
				loads.Set(h, 0, p.nodeForce.M)
			}
		}
	}

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
				stiffinerKGlobal.Set(g, h, 0.0)
				stiffinerKGlobal.Set(h, g, 0.0)
			}
			stiffinerKGlobal.Set(g, g, 1.0)
			// modify load vector on support
			loads.Set(g, 0, 0.0)
		}
	}

	{
		// If all elements of loads is zero,
		// then somethink wrong
		var haveNonZeroElements bool
		eps := 1e-8
		for i := 0; i < loads.GetRowSize(); i++ {
			if loads.Get(i, 0) > eps || loads.Get(i, 0) < -eps {
				haveNonZeroElements = true
				break
			}
		}
		if !haveNonZeroElements {
			return fmt.Errorf("Vector of loads have only zero elements")
		}
	}

	// Solving system of linear equations for finding
	// the displacement in points in global system
	// TODO: one global stiffiner matrix for all cases
	lu := solver.NewLUsolver(stiffinerKGlobal)
	globalDisp := lu.Solve(loads)

	// global displacement for points
	for _, p := range m.points {
		axes := m.degreeForPoint.GetDoF(p.Index)
		var disp displacement.PointDim2
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
	}

	for _, ele := range m.elements {
		switch ele.(type) {
		case element.Beam:
			beam := ele.(element.Beam)
			fe := m.getBeamFiniteElement(beam.GetIndex())
			_, degreeLocal := finiteElement.GetStiffinerGlobalK(fe, &m.degreeForPoint, finiteElement.FullInformation)
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
			// TODO : only for 2d
			forceCase.localDisplacement = append(forceCase.localDisplacement, displacement.BeamDim2{
				Begin: displacement.Dim2{
					Dx: localDisplacement[0],
					Dy: localDisplacement[1],
					Dm: localDisplacement[2],
				},
				End: displacement.Dim2{
					Dx: localDisplacement[3],
					Dy: localDisplacement[4],
					Dm: localDisplacement[5],
				},
				Index: beam.GetIndex(),
			})

			klocal := matrix.NewMatrix64bySize(10, 10)
			fe.GetStiffinerK(&klocal)

			var localForce []float64
			for i := 0; i < klocal.GetRowSize(); i++ {
				sum := 0.0
				for j := 0; j < klocal.GetRowSize(); j++ {
					sum += klocal.Get(i, j) * localDisplacement[j]
				}
				localForce = append(localForce, sum)
			}
			// TODO : only for 2d
			forceCase.localForces = append(forceCase.localForces, forceLocal.BeamDim2{
				Begin: forceLocal.Dim2{
					Fx: localForce[0],
					Fy: localForce[1],
					M:  localForce[2],
				},
				End: forceLocal.Dim2{
					Fx: localForce[3],
					Fy: localForce[4],
					M:  localForce[5],
				},
				Index: beam.GetIndex(),
			})
		default:
			panic("")
		}
	}

	// reactions
	for _, r := range m.supports {
		var reac reaction.Dim2
		reac.Index = r.pointIndex
		for _, e := range m.elements {
			switch e.(type) {
			case element.Beam:
				beam := e.(element.Beam)
				if beam.GetPointIndex()[0] != r.pointIndex && beam.GetPointIndex()[1] != r.pointIndex {
					continue
				}
				// KG get matrix stiffiner for beam in global system
				fe := m.getBeamFiniteElement(beam.GetIndex())
				k, _ := finiteElement.GetStiffinerGlobalK(fe, &m.degreeForPoint, finiteElement.FullInformation)
				// D get vector displacement in global system
				// TODO : only for 2d
				d := matrix.NewMatrix64bySize(6, 1)
				{
					g, err := forceCase.GetGlobalDisplacement(beam.GetPointIndex()[0])
					if err != nil {
						return err
					}
					// TODO : only for 2d
					d.Set(0, 0, g.Dx)
					d.Set(1, 0, g.Dy)
					d.Set(2, 0, g.Dm)
				}
				{
					g, err := forceCase.GetGlobalDisplacement(beam.GetPointIndex()[1])
					if err != nil {
						return err
					}
					// TODO : only for 2d
					d.Set(3, 0, g.Dx)
					d.Set(4, 0, g.Dy)
					d.Set(5, 0, g.Dm)
				}
				// L = KG*D, L - loads in global system
				L := k.Times(&d)
				if beam.GetPointIndex()[0] == r.pointIndex {
					// sum in reaction
					// TODO : only for 2d
					reac.Fx += L.Get(0, 0)
					reac.Fy += L.Get(1, 0)
					reac.M += L.Get(2, 0)
				}
				if beam.GetPointIndex()[1] == r.pointIndex {
					// sum in reaction
					// TODO : only for 2d
					reac.Fx += L.Get(3, 0)
					reac.Fy += L.Get(4, 0)
					reac.M += L.Get(5, 0)
				}
			default:
				panic("please add finite element in reactions")
			}
		}
		forceCase.reactions = append(forceCase.reactions, reac)
	}

	return nil
}
