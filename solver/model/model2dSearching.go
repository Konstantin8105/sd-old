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
func (m *Dim2) getMaterial(index element.Index) (mat material.Linear, err error) {
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
func (m *Dim2) getCoordinate(index element.Index) (c []point.Dim2, err error) {
	var inx []point.Index
	var found bool
	for _, e := range m.elements {
		switch e.(type) {
		case element.Beam:
			beam := e.(element.Beam)
			if beam.GetIndex() != index {
				continue
			}
			for i := range beam.GetPointIndex() {
				inx = append(inx, beam.GetPointIndex()[i])
			}
			found = true
			break

		default:
			panic("")
		}
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
	for _, group := range m.truss {
		for _, inx := range group.beamIndexes {
			if inx == index {
				return true
			}
		}
	}
	return false
}
