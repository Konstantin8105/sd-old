package model

import (
	"fmt"

	"github.com/Konstantin8105/GoFea/input/element"
	"github.com/Konstantin8105/GoFea/input/point"
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

	// check points
	err := isUniqueIndexes(pointsByPoints(m.points))
	if err != nil {
		return fmt.Errorf("Errors in poins:\n%v", err)
	}

	//check elements
	err = isUniqueIndexes(elementsByGetIndex(m.elements))
	if err != nil {
		return fmt.Errorf("Errors in elements:\n%v", err)
	}

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
		err := f.check()
		if err != nil {
			return err
		}
	}

	// compress material
	err = isUniqueIndexes(materialByElement(m.materials))
	if err != nil {
		return fmt.Errorf("Errors in material:\n%v", err)
	}

	// compress support
	err = isUniqueIndexes(supportByPoint(m.supports))
	if err != nil {
		return fmt.Errorf("Errors in supports:\n%v", err)
	}

	// compress shape
	err = isUniqueIndexes(shapeByElement(m.shapes))
	if err != nil {
		return fmt.Errorf("Errors in shape:\n%v", err)
	}

	// compress truss
	err = isUniqueIndexes(elementsByElements(m.truss))
	if err != nil {
		return fmt.Errorf("Errors in truss:\n%v", err)
	}

	//TODO  Example of use : sort.Sort(materialByElement(slise))
	//TODO sorting for quick search - quick checking

	return nil
}

type elementsByGetIndex []element.Elementer

func (a elementsByGetIndex) Len() int            { return len(a) }
func (a elementsByGetIndex) Swap(i, j int)       { a[i], a[j] = a[j], a[i] }
func (a elementsByGetIndex) Less(i, j int) bool  { return a[i].GetIndex() < a[j].GetIndex() }
func (a elementsByGetIndex) Equal(i, j int) bool { return a[i].GetIndex() == a[j].GetIndex() }
func (a elementsByGetIndex) Name(i int) int      { return int(a[i].GetIndex()) }

type elementsByElements []element.Index

func (a elementsByElements) Len() int            { return len(a) }
func (a elementsByElements) Swap(i, j int)       { a[i], a[j] = a[j], a[i] }
func (a elementsByElements) Less(i, j int) bool  { return a[i] < a[j] }
func (a elementsByElements) Equal(i, j int) bool { return a[i] == a[j] }
func (a elementsByElements) Name(i int) int      { return int(a[i]) }

type pointsByPoints []point.Dim2

func (a pointsByPoints) Len() int            { return len(a) }
func (a pointsByPoints) Swap(i, j int)       { a[i], a[j] = a[j], a[i] }
func (a pointsByPoints) Less(i, j int) bool  { return a[i].Index < a[j].Index }
func (a pointsByPoints) Equal(i, j int) bool { return a[i].Index == a[j].Index }
func (a pointsByPoints) Name(i int) int      { return int(a[i].Index) }

type uniqueElements interface {
	Len() int
	Equal(i, j int) bool
	Name(i int) int
}

func isUniqueIndexes(u uniqueElements) error {
	size := u.Len()
	var errorIndexes []int
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			if u.Equal(i, j) {
				errorIndexes = append(errorIndexes, i)
				errorIndexes = append(errorIndexes, j)
			}
		}
	}
	if len(errorIndexes) > 0 {
		s := "Please clarify, because next elements or points have same index:\n"
		for i := 0; i < len(errorIndexes); i += 2 {
			s += fmt.Sprintf("%v and %v\n", u.Name(errorIndexes[i]), u.Name(errorIndexes[i+1]))
		}
		return fmt.Errorf("Error. %v", s)
	}
	return nil
}
