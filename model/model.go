package model

import (
	coordinate "github.com/Konstantin8105/GoFea/Coordinate"
	element "github.com/Konstantin8105/GoFea/Element"
	force "github.com/Konstantin8105/GoFea/Force"
	material "github.com/Konstantin8105/GoFea/Material"
	shape "github.com/Konstantin8105/GoFea/Shape"
	support "github.com/Konstantin8105/GoFea/Support"
)

type supportStruct struct {
	support         support.Support
	coordinateIndex []coordinate.PointIndex
}

type shapeStruct struct {
	shape     shape.Shape
	beamIndex []element.BeamIndex
}

type materialStruct struct {
	material  material.Material
	beamIndex []element.BeamIndex
}

type gravityForceStruct struct {
	gravityForce force.GravityForce
	beamIndex    []element.BeamIndex
}

type nodeForceStruct struct {
	nodeForce       force.NodeForce
	coordinateIndex []coordinate.PointIndex
}

type forceCase struct {
	gravityForces []gravityForceStruct
	nodeForces    []nodeForceStruct
}

// Model - base struct of data for model
type Model struct {
	coordinates []coordinate.Coordinate
	beams       []element.Beam
	supports    []supportStruct
	shapes      []shapeStruct
	materials   []materialStruct
	forceCases  forceCase
}

// AddCoodinates - add coordinate to model
func (m *Model) AddCoodinates(coordinates ...coordinate.Coordinate) {
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
func (m *Model) AddMaterial(material material.Material, beams ...element.BeamIndex) {
	m.materials = append(m.materials, materialStruct{
		material:  material,
		beamIndex: beams,
	})
}

// AddSupport - add support for points
func (m *Model) AddSupport(support support.Support, coordinateIndex ...coordinate.PointIndex) {
	m.supports = append(m.supports, supportStruct{
		support:         support,
		coordinateIndex: coordinateIndex,
	})
}

// AddGravityForce - add gravity force
func (m *Model) AddGravityForce(g force.GravityForce, beams ...element.BeamIndex) {
	m.forceCases.gravityForces = append(m.forceCases.gravityForces, gravityForceStruct{
		gravityForce: g,
		beamIndex:    beams,
	})
}

// AddNodeForce - add node force
func (m *Model) AddNodeForce(n force.NodeForce, points ...coordinate.PointIndex) {
	m.forceCases.nodeForces = append(m.forceCases.nodeForces, nodeForceStruct{
		nodeForce:       n,
		coordinateIndex: points,
	})
}
