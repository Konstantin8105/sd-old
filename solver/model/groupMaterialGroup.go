package model

import (
	"github.com/Konstantin8105/GoFea/input/element"
	"github.com/Konstantin8105/GoFea/input/material"
)

type materialLinearGroup struct {
	material     material.Linear
	elementIndex element.Index
}

// Sorting
type materialByElement []materialLinearGroup

func (a materialByElement) Len() int            { return len(a) }
func (a materialByElement) Swap(i, j int)       { a[i], a[j] = a[j], a[i] }
func (a materialByElement) Less(i, j int) bool  { return a[i].elementIndex < a[j].elementIndex }
func (a materialByElement) Equal(i, j int) bool { return a[i].elementIndex == a[j].elementIndex }
func (a materialByElement) Name(i int) int      { return int(a[i].elementIndex) }
