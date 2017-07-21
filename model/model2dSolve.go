package model

import (
	"fmt"
	"sync"

	"github.com/Konstantin8105/GoFea/dof"
	"github.com/Konstantin8105/GoFea/element"
	"github.com/Konstantin8105/GoFea/finiteElement"
	"github.com/Konstantin8105/GoLinAlg/matrix"
)

// Solve - solving finite element
func (m *Dim2) Solve() (err error) {

	// create error channel
	type results struct {
		err       error
		forceCase *forceCase2d
	}

	// summary result
	var summaryResult []results

	resCh := make(chan results)
	go func() {
		for rc := range resCh {
			summaryResult = append(summaryResult, rc)
		}
	}()

	// create workgroup
	var wg sync.WaitGroup

	for i := range m.forceCases {
		wg.Add(1)
		go func(f *forceCase2d) {
			// work is done
			defer wg.Done()

			err := m.solveCase(f)
			resCh <- results{
				err:       err,
				forceCase: f,
			}
		}(&(m.forceCases[i]))
	}
	wg.Wait()

	close(resCh)

	var haveError bool
	for _, s := range summaryResult {
		if s.err != nil {
			haveError = true
			break
		}
	}
	if !haveError {
		return nil
	}

	// TODO: more beatiful
	return fmt.Errorf("%#v", summaryResult)
}

func (m *Dim2) getBeamFiniteElement(inx element.BeamIndex) (fe finiteElement.FiniteElementer) {
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
		f := finiteElement.TrussDim2{
			Material: material,
			Shape:    shape,
			Points:   coord,
		}
		return &f
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
	return nil
}

func (m *Dim2) convertFromLocalToGlobalSystem(degreeGlobal *[]dof.AxeNumber, dofSystem *dof.DoF, mapIndex *dof.MapIndex, f func(finiteElement.FiniteElementer, *dof.DoF, finiteElement.Information) (matrix.T64, []dof.AxeNumber)) matrix.T64 {
	globalResult := matrix.NewMatrix64bySize(len(*degreeGlobal), len(*degreeGlobal))
	for _, beam := range m.beams {
		fe := m.getBeamFiniteElement(beam.Index)
		klocal, degreeLocal := f(fe, dofSystem, finiteElement.WithoutZeroStiffiner)
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
				globalResult.Set(g, h, globalResult.Get(g, h)+klocal.Get(i, j))
			}
		}
	}
	return globalResult
}
