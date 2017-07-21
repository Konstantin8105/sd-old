package model

import (
	"fmt"

	"github.com/Konstantin8105/GoFea/element"
	"github.com/Konstantin8105/GoFea/material"
	"github.com/Konstantin8105/GoFea/point"
	"github.com/Konstantin8105/GoFea/shape"
)

// getShape - searching shape for beam
func (m *Dim2) getShape(index element.ElementIndex) (s shape.Shape, err error) {
	for _, shapes := range m.shapes {
		for _, inx := range shapes.beamIndexes {
			if inx == index {
				return shapes.shape, nil
			}
		}
	}
	return s, fmt.Errorf("Cannot found shape for beam #%v", index)
}

// getMaterial - searching material for beam
func (m *Dim2) getMaterial(index element.ElementIndex) (mat material.Linear, err error) {
	for _, material := range m.materials {
		for _, inx := range material.beamIndexes {
			if index == inx {
				mat = material.material
				err = nil
				return
			}
		}
	}
	return mat, fmt.Errorf("Cannot found material for beam #%v", index)
}

// getCoordinate - return coordinate of beam
func (m *Dim2) getCoordinate(index element.ElementIndex) (c [2]point.Dim2, err error) {
	var inx [2]point.Index
	var found bool
	for _, beam := range m.elements {
		if beam.Index == index {
			inx = beam.PointIndexes
			found = true
			break
		}
	}
	if !found {
		return c, fmt.Errorf("Cannot found beam with index #%v", index)
	}

	var coord [2]point.Dim2
	for i := 0; i < 2; i++ {
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
func (m *Dim2) isTruss(index element.ElementIndex) bool {
	for _, group := range m.truss {
		for _, inx := range group.beamIndexes {
			if inx == index {
				return true
			}
		}
	}
	return false
}
