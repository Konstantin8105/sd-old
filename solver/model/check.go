package model

import (
	"fmt"

	"github.com/Konstantin8105/GoFea/input/element"
	"github.com/Konstantin8105/GoFea/utils"
)

func (m *Dim2) checkInputData() error {
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

	// checking lenght of finite element beam
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
		return fmt.Errorf("Finite element %s have lenght equal zero", list)
	}

	// checking beam with same number

	//

	return nil
}
