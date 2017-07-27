package element

import "github.com/Konstantin8105/GoFea/input/point"

var _ Elementer = (*Beam)(nil)

// Elementer - interface typical for elements
type Elementer interface {
	// GetIndex - return index of element
	GetIndex() Index

	// GetPointIndex - return indexes of point for that finite element
	GetPointIndex() []point.Index

	// GetAmountPoint - return amount points in finite element
	GetAmountPoint() int
}
