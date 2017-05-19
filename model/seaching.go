package model

import (
	"fmt"

	"github.com/Konstantin8105/GoFea/element"
	"github.com/Konstantin8105/GoFea/material"
	"github.com/Konstantin8105/GoFea/point"
	"github.com/Konstantin8105/GoFea/shape"
)

// GetShape - searching shape for beam
func (m *Model) GetShape(index element.BeamIndex) (s shape.Shape, err error) {
	for _, shapes := range m.shapes {
		for _, inx := range shapes.beamIndex {
			if inx == index {
				return shapes.shape, nil
			}
		}
	}
	return s, fmt.Errorf("Cannot found shape for beam #%v", index)
}

// GetMaterial - searching material for beam
func (m *Model) GetMaterial(index element.BeamIndex) (mat material.Linear, err error) {
	for _, material := range m.materials {
		for _, inx := range material.beamIndex {
			if index == inx {
				mat = material.material
				err = nil
				return
			}
		}
	}
	return mat, fmt.Errorf("Cannot found material for beam #%v", index)
}

// GetCoordinate - return coordinate of beam
func (m *Model) GetCoordinate(index element.BeamIndex) (c [2]point.Dim3, err error) {
	var inx [2]point.Index
	var found bool
	for _, beam := range m.beams {
		if beam.Index == index {
			inx = beam.PointIndexs
			found = true
			break
		}
	}
	if !found {
		return c, fmt.Errorf("Cannot found beam with index #%v", index)
	}

	var coord [2]point.Dim3
	for i := 0; i < 2; i++ {
		found = false
		for _, c := range m.coordinates {
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

/*
//GetLenght - calculate lenght of beam
func (m *Model) GetLenght(index element.BeamIndex) (lenght float64, err error) {
	coord, err := m.GetCoordinate(index)
	if err != nil {
		return 0.0, err
	}
	return math.Sqrt(math.Pow(coord[0].X-coord[1].X, 2.0) + math.Pow(coord[0].Y-coord[1].Y, 2.0) + math.Pow(coord[0].Z-coord[1].Z, 2.0)), nil
}
*/
