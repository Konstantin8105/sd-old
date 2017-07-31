package model

import (
	"github.com/Konstantin8105/GoFea/input/element"
	"github.com/Konstantin8105/GoFea/input/shape"
)

type shapeGroup struct {
	shape        shape.Shape
	elementIndex element.Index
}

// Sorting
type shapeByElement []shapeGroup

func (a shapeByElement) Len() int            { return len(a) }
func (a shapeByElement) Swap(i, j int)       { a[i], a[j] = a[j], a[i] }
func (a shapeByElement) Less(i, j int) bool  { return a[i].elementIndex < a[j].elementIndex }
func (a shapeByElement) Equal(i, j int) bool { return a[i].elementIndex == a[j].elementIndex }
func (a shapeByElement) Name(i int) int      { return int(a[i].elementIndex) }
