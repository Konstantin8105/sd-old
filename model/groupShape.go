package model

import (
	"github.com/Konstantin8105/GoFea/element"
	"github.com/Konstantin8105/GoFea/shape"
)

type shapeGroup struct {
	shape       shape.Shape
	beamIndexes []element.ElementIndex
}
