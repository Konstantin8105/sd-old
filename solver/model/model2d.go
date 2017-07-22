package model

import (
	"github.com/Konstantin8105/GoFea/input/element"
	"github.com/Konstantin8105/GoFea/input/force"
	"github.com/Konstantin8105/GoFea/input/material"
	"github.com/Konstantin8105/GoFea/input/point"
	"github.com/Konstantin8105/GoFea/input/shape"
	"github.com/Konstantin8105/GoFea/input/support"
)

// Dim2 - base struct of data for model in 2d
type Dim2 struct {
	// input data
	points     []point.Dim2
	elements   []element.Elementer
	truss      []trussGroup
	supports   []supportGroup2d
	shapes     []shapeGroup
	materials  []materialLinearGroup
	forceCases []forceCase2d
}

// AddPoint - add point to model
func (m *Dim2) AddPoint(points ...point.Dim2) {
	m.points = append(m.points, points...)
}

// AddElement - add beam to model
func (m *Dim2) AddElement(elements ...element.Elementer) {
	m.elements = append(m.elements, elements...)
}

// AddTrussProperty - add truss property for beam
func (m *Dim2) AddTrussProperty(beamIndexes ...element.ElementIndex) {
	m.truss = append(m.truss, trussGroup{beamIndexes: beamIndexes})
}

// AddSupport - add support for points
func (m *Dim2) AddSupport(support support.Dim2, pointIndexes ...point.Index) {
	m.supports = append(m.supports, supportGroup2d{
		support:      support,
		pointIndexes: pointIndexes,
	})
}

// AddShape - add shape property for beam
func (m *Dim2) AddShape(shape shape.Shape, beamIndexes ...element.ElementIndex) {
	m.shapes = append(m.shapes, shapeGroup{
		shape:       shape,
		beamIndexes: beamIndexes,
	})
}

// AddMaterial - add material for beam
func (m *Dim2) AddMaterial(material material.Linear, beamIndexes ...element.ElementIndex) {
	m.materials = append(m.materials, materialLinearGroup{
		material:    material,
		beamIndexes: beamIndexes,
	})
}

// AddNodeForce - add node force in force case
func (m *Dim2) AddNodeForce(caseNumber int, nodeForce force.NodeDim2, pointIndexes ...point.Index) {
	for i := range m.forceCases {
		if m.forceCases[i].indexCase == caseNumber {
			m.forceCases[i].nodeForces = append(m.forceCases[i].nodeForces, nodeForce2d{
				nodeForce:    nodeForce,
				pointIndexes: pointIndexes,
			})
			return
		}
	}

	nf := nodeForce2d{
		nodeForce:    nodeForce,
		pointIndexes: pointIndexes,
	}

	var fc forceCase2d
	fc.indexCase = caseNumber
	fc.nodeForces = append(fc.nodeForces, nf)

	m.forceCases = append(m.forceCases, fc)
}

// AddGravityForce - add gravity force in force case
func (m *Dim2) AddGravityForce(caseNumber int, gravityForce force.GravityDim2, beamIndexes ...element.ElementIndex) {
	for i := range m.forceCases {
		if m.forceCases[i].indexCase == caseNumber {
			m.forceCases[i].gravityForces = append(m.forceCases[i].gravityForces, gravityForce2d{
				gravityForce: gravityForce,
				beamIndexes:  beamIndexes,
			})
			return
		}
	}

	gf := gravityForce2d{
		gravityForce: gravityForce,
		beamIndexes:  beamIndexes,
	}

	var fc forceCase2d
	fc.indexCase = caseNumber
	fc.gravityForces = append(fc.gravityForces, gf)

	m.forceCases = append(m.forceCases, fc)
}
