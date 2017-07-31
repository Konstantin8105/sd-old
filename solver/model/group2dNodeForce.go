package model

import (
	"github.com/Konstantin8105/GoFea/input/force"
	"github.com/Konstantin8105/GoFea/input/point"
)

type nodeForce2d struct {
	nodeForce  force.NodeDim2
	pointIndex point.Index
}

// Sorting
type nodeForceByPoint []nodeForce2d

func (a nodeForceByPoint) Len() int { return len(a) }

//func (a nodeForceByPoint) Swap(i, j int)       { a[i], a[j] = a[j], a[i] }
//func (a nodeForceByPoint) Less(i, j int) bool  { return a[i].pointIndex < a[j].pointIndex }
func (a nodeForceByPoint) Equal(i, j int) bool { return a[i].pointIndex == a[j].pointIndex }
func (a nodeForceByPoint) Name(i int) int      { return int(a[i].pointIndex) }
