package model

import (
	"github.com/Konstantin8105/GoFea/element"
	"github.com/Konstantin8105/GoFea/force"
	"github.com/Konstantin8105/GoFea/material"
	"github.com/Konstantin8105/GoFea/point"
	"github.com/Konstantin8105/GoFea/shape"
	"github.com/Konstantin8105/GoFea/support"
)

type supportStruct struct {
	support         support.Dim3
	coordinateIndex []point.Index
}

type shapeStruct struct {
	shape     shape.Shape
	beamIndex []element.BeamIndex
}

type materialStruct struct {
	material  material.Linear
	beamIndex []element.BeamIndex
}

type gravityForceStruct struct {
	gravityForce force.GravityDim3
	beamIndex    []element.BeamIndex
}

type nodeForceStruct struct {
	nodeForce       force.NodeDim3
	coordinateIndex []point.Index
}

type forceCase struct {
	gravityForces []gravityForceStruct
	nodeForces    []nodeForceStruct
}

// Model - base struct of data for model
type Model struct {
	coordinates []point.Dim3
	beams       []element.Beam
	supports    []supportStruct
	shapes      []shapeStruct
	materials   []materialStruct
	forceCases  forceCase
}

// AddCoodinates - add coordinate to model
func (m *Model) AddCoodinates(coordinates ...point.Dim3) {
	m.coordinates = append(m.coordinates, coordinates...)
}

// AddBeams - add beam to model
func (m *Model) AddBeams(beams ...element.Beam) {
	m.beams = append(m.beams, beams...)
}

// AddShape - add shape property for beam
func (m *Model) AddShape(shape shape.Shape, beams ...element.BeamIndex) {
	m.shapes = append(m.shapes, shapeStruct{
		shape:     shape,
		beamIndex: beams,
	})
}

// AddMaterial - add meaterial for beam
func (m *Model) AddMaterial(material material.Linear, beams ...element.BeamIndex) {
	m.materials = append(m.materials, materialStruct{
		material:  material,
		beamIndex: beams,
	})
}

// AddSupport - add support for points
func (m *Model) AddSupport(support support.Dim3, coordinateIndex ...point.Index) {
	m.supports = append(m.supports, supportStruct{
		support:         support,
		coordinateIndex: coordinateIndex,
	})
}

// AddGravityForce - add gravity force
func (m *Model) AddGravityForce(g force.GravityDim3, beams ...element.BeamIndex) {
	m.forceCases.gravityForces = append(m.forceCases.gravityForces, gravityForceStruct{
		gravityForce: g,
		beamIndex:    beams,
	})
}

// AddNodeForce - add node force
func (m *Model) AddNodeForce(n force.NodeDim3, points ...point.Index) {
	m.forceCases.nodeForces = append(m.forceCases.nodeForces, nodeForceStruct{
		nodeForce:       n,
		coordinateIndex: points,
	})
}
