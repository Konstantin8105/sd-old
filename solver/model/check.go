package model

import "fmt"

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

	// checking two point with same number

	// checking beam with same number

	// checking lenght of finite element beam
	// for avoid divide by zero

	return nil
}
