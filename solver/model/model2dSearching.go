package model

import (
	"fmt"

	"github.com/Konstantin8105/GoFea/input/element"
	"github.com/Konstantin8105/GoFea/input/material"
	"github.com/Konstantin8105/GoFea/input/point"
	"github.com/Konstantin8105/GoFea/input/shape"
)

// getShape - searching shape for beam
func (m *Dim2) getShape(index element.Index) (s shape.Shape, err error) {
	//TODO change to binary search
	for _, sh := range m.shapes {
		if sh.elementIndex == index {
			return sh.shape, nil
		}
	}
	return s, fmt.Errorf("Cannot found shape for beam #%v", index)
}

// getMaterial - searching material for beam
func (m *Dim2) getMaterial(index element.Index) (mat material.Linear, err error) {
	//TODO change to binary search
	for _, m := range m.materials {
		if m.elementIndex == index {
			return m.material, nil
		}
	}
	return mat, fmt.Errorf("Cannot found material for beam #%v", index)
}

// getCoordinate - return coordinate of beam
func (m *Dim2) getCoordinate(index element.Index) (c []point.Dim2, err error) {
	//TODO change to binary search
	var inx []point.Index
	var found bool
	for _, e := range m.elements {
		if e.GetIndex() != index {
			continue
		}
		for i := range e.GetPointIndex() {
			inx = append(inx, e.GetPointIndex()[i])
		}
		found = true
		break
	}
	if !found {
		return c, fmt.Errorf("Cannot found beam with index #%v", index)
	}

	coord := make([]point.Dim2, len(inx))
	for i := 0; i < len(inx); i++ {
		found = false
		for _, c := range m.points {
			if inx[i] == c.Index {
				found = true
				coord[i] = c
				break
			}
		}
		if !found {
			return c, fmt.Errorf("Cannot found coordinate for beam index #%v", i)
		}
	}
	return coord, nil
}

// isTruss - return true if beam is truss
func (m *Dim2) isTruss(index element.Index) bool {
	//TODO change to binary search
	for _, t := range m.truss {
		if t == index {
			return true
		}
	}
	return false
}

func (m *Dim2) getElement(index element.Index) (element.Elementer, error) {
	//TODO change to binary search
	for _, t := range m.elements {
		if t.GetIndex() == index {
			return t, nil
		}
	}
	return nil, fmt.Errorf("cannot found element %v", index)
}
