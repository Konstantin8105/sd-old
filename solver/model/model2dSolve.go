package model

import (
	"fmt"

	"github.com/Konstantin8105/GoFea/input/element"
	"github.com/Konstantin8105/GoFea/input/point"
	"github.com/Konstantin8105/GoFea/solver/dof"
	"github.com/Konstantin8105/GoFea/solver/finiteElement"
	"github.com/Konstantin8105/GoLinAlg/matrix"
)

// Solve - solving finite element
func (m *Dim2) Solve() (err error) {

	err = m.checkInputData()
	if err != nil {
		return err
	}

	// generate degrees of freedom
	err = m.generateDof()
	if err != nil {
		return err
	}

	// TODO : check everything
	// TODO : sort  everything
	// TODO : compress loads by number

	// create error channel
	type results struct {
		err       error
		forceCase int
	}

	// summary result
	var summaryResult []results

	for i := range m.forceCases {
		err := m.solveCase(&(m.forceCases[i]))
		summaryResult = append(summaryResult, results{
			err:       err,
			forceCase: m.forceCases[i].indexCase,
		})
		switch m.forceCases[i].dynamicType {
		case naturalFrequency:
			err := m.solveNaturalFrequency(&(m.forceCases[i]))
			summaryResult = append(summaryResult, results{
				err:       err,
				forceCase: m.forceCases[i].indexCase,
			})
		case bucklingFactors:
			panic("add buckling algorithm")
		}
	}

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

	// TODO: more beautiful
	return fmt.Errorf("%#v", summaryResult)
}

func (m *Dim2) getFiniteElement(inx element.Index) (fe finiteElement.FiniteElementer, err error) {
	material, err := m.getMaterial(inx)
	if err != nil {
		return fe, fmt.Errorf("Cannot found material for beam #%v. Error = %v", inx, err)
	}
	shape, err := m.getShape(inx)
	if err != nil {
		return fe, fmt.Errorf("Cannot found shape for beam #%v. Error = %v", inx, err)
	}
	el, err := m.getElement(inx)
	if err != nil {
		return fe, fmt.Errorf("Cannot found element %v. Error = %v", inx, err)
	}
	coord, err := m.getCoordinate(inx)
	if err != nil {
		return fe, fmt.Errorf("Cannot calculate length for beam #%v. Error = %v", inx, err)
	}

	switch el.(type) {
	case element.Beam:
		if m.isTruss(inx) {
			// No need to check on len(points) == 2
			var c [2]point.Dim2
			for i := 0; i < len(coord); i++ {
				c[i] = coord[i]
			}
			f := finiteElement.TrussDim2{
				Material: material,
				Shape:    shape,
				Points:   c,
			}
			return &f, nil
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
	}
	panic("Please add finite element")
}

func (m *Dim2) convertFromLocalToGlobalSystem(degreeGlobal *[]dof.AxeNumber, dofSystem *dof.DoF, mapIndex *dof.MapIndex, f func(finiteElement.FiniteElementer, *dof.DoF, finiteElement.Information) (matrix.T64, []dof.AxeNumber)) (y matrix.T64, err error) {
	globalResult := matrix.NewMatrix64bySize(len(*degreeGlobal), len(*degreeGlobal))
	for _, ele := range m.elements {
		fe, err := m.getFiniteElement(ele.GetIndex())
		if err != nil {
			return y, err
		}
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
	return globalResult, nil
}
