package model

import (
	"github.com/Konstantin8105/GoFea/input/element"
	"github.com/Konstantin8105/GoFea/input/shape"
)

type shapeGroup struct {
	shape          shape.Shape
	elementIndexes element.Index
}
