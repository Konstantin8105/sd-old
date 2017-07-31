package model

import (
	"github.com/Konstantin8105/GoFea/input/point"
	"github.com/Konstantin8105/GoFea/input/support"
)

type supportGroup2d struct {
	support    support.Dim2
	pointIndex point.Index
}

// Sorting
type supportByPoint []supportGroup2d

func (a supportByPoint) Len() int            { return len(a) }
func (a supportByPoint) Swap(i, j int)       { a[i], a[j] = a[j], a[i] }
func (a supportByPoint) Less(i, j int) bool  { return a[i].pointIndex < a[j].pointIndex }
func (a supportByPoint) Equal(i, j int) bool { return a[i].pointIndex == a[j].pointIndex }
func (a supportByPoint) Name(i int) int      { return int(a[i].pointIndex) }
