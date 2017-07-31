package model

import (
	"fmt"

	"github.com/Konstantin8105/GoFea/input/element"
	"github.com/Konstantin8105/GoFea/utils"
)

func (m *Dim2) checkInputData() error {
	// TODO add common error slise

	errorText := "Not enought data for calculate. %s"
	if len(m.points) < 2 {
		return fmt.Errorf(errorText, "Please add points in model")
	}
	if len(m.elements) < 1 {
		return fmt.Errorf(errorText, "Please add finite elements in model")
	}
	if len(m.supports) < 1 {
		return fmt.Errorf(errorText, "Please add supports in model")
	}
	if len(m.materials) < 1 {
		return fmt.Errorf(errorText, "Please add material in model")
	}
	if len(m.forceCases) < 1 {
		return fmt.Errorf(errorText, "Please add load case in model")
	}

	// checking "amount of point in finite element" for example:
	// - only 2 points in beam
	// - ...
	// no need

	sorting for quick search - quick checking

	// checking length of finite element beam
	// for avoid divide by zero
	var zeroElements []element.Index
	for _, e := range m.elements {
		coord, err := m.getCoordinate(e.GetIndex())
		if err != nil {
			return err
		}
		for i := 0; i < len(coord); i++ {
			for j := 0; j < len(coord); j++ {
				if i <= j {
					continue
				}
				if utils.LengthDim2(coord[i], coord[j]) <= 0.0 {
					zeroElements = append(zeroElements, e.GetIndex())
					goto next
				}
			}
		}
	next:
	}
	if len(zeroElements) > 0 {
		var list string
		for i := range zeroElements {
			list += fmt.Sprintf("%v,", zeroElements[i])
		}
		return fmt.Errorf("Finite element %s have length equal zero", list)
	}

	// compress node loads
	for _, f := range m.forceCases {
	begin:
		size := len(f.nodeForces)
		for i := 0; i < size; i++ {
			for j := i + 1; j < size; j++ {
				if f.nodeForces[i].pointIndex == f.nodeForces[j].pointIndex {
					another algorithm - single node for single point
					f.nodeForces[i].nodeForce.Plus(f.nodeForces[j].nodeForce)
					for k := j; k < size-1; k++ {
						f.nodeForces[k] = f.nodeForces[k+1]
					}
					f.nodeForces = f.nodeForces[0 : len(f.nodeForces)-1]
					goto begin
				}
			}
		}
	}

	// compress material
	{
		size := len(m.materials)
		var errorIndexes []int
		for i := 0; i < size; i++ {
			for j := i + 1; j < size; j++ {
				if m.materials[i].elementIndex == m.materials[j].elementIndex {
					errorIndexes = append(errorIndexes, i)
					errorIndexes = append(errorIndexes, j)
				}
			}
		}
		if len(errorIndexes) > 0 {
			s := "Please clarify material, because material is same for next elements:\n"
			for i := 0; i < len(errorIndexes); i += 2 {
				s += fmt.Sprintf("%v and %v\n", errorIndexes[i], errorIndexes[i+1])
			}
			return fmt.Errorf("Error. %v", s)
		}
	}

	// compress support
	{
	beginSupport:
		size := len(m.supports)
		for i := 0; i < size; i++ {
			for j := i + 1; j < size; j++ {
				if m.supports[i].pointIndex == m.supports[j].pointIndex {
					for k := j; k < size-1; k++ {
						m.supports[k] = m.supports[k+1]
					}
					m.supports = m.supports[0 : len(m.supports)-1]
					goto beginSupport
				}
			}
		}
	}

	// compress shape
	{
		size := len(m.shapes)
		var errorIndexes []int
		for i := 0; i < size; i++ {
			for j := i + 1; j < size; j++ {
				if m.shapes[i].elementIndex == m.shapes[j].elementIndex {
					errorIndexes = append(errorIndexes, i)
					errorIndexes = append(errorIndexes, j)
				}
			}
		}
		if len(errorIndexes) > 0 {
			s := "Please clarify shape, because shapes is same for next elements:\n"
			for i := 0; i < len(errorIndexes); i += 2 {
				s += fmt.Sprintf("%v and %v\n", errorIndexes[i], errorIndexes[i+1])
			}
			return fmt.Errorf("Error. %v", s)
		}
	}

	// compress truss
	{
	beginTruss:
		size := len(m.truss)
		for i := 0; i < size; i++ {
			for j := i + 1; j < size; j++ {
				if m.truss[i] == m.truss[j] {
					another algorithm - single truss for single element
					for k := j; k < size-1; k++ {
						m.truss[k] = m.truss[k+1]
					}
					m.truss = m.truss[0 : len(m.truss)-1]
					goto beginTruss
				}
			}
		}
	}


	return nil
}
