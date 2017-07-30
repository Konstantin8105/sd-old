package model

import (
	"fmt"

	"github.com/Konstantin8105/GoFea/input/element"
	"github.com/Konstantin8105/GoFea/input/force"
	"github.com/Konstantin8105/GoFea/input/material"
	"github.com/Konstantin8105/GoFea/input/point"
	"github.com/Konstantin8105/GoFea/input/shape"
	"github.com/Konstantin8105/GoFea/input/support"
	"github.com/Konstantin8105/GoFea/output/displacement"
	"github.com/Konstantin8105/GoFea/output/forceLocal"
	"github.com/Konstantin8105/GoFea/output/reaction"
	"github.com/Konstantin8105/GoFea/solver/dof"
)

// Dim2 - base struct of data for model in 2d
type Dim2 struct {
	// input data
	points     []point.Dim2
	elements   []element.Elementer
	truss      []element.Index
	supports   []supportGroup2d
	shapes     []shapeGroup
	materials  []materialLinearGroup
	forceCases []forceCase2d

	//TODO: if we have nolinear finite element then,
	// dof can be different between load cases

	// internal data
	degreeInGlobalMatrix []dof.AxeNumber // degree of freedom in global system, created in according to "real" finite elements and it is not the same "dofSystem" for models with many pin.
	indexsInGlobalMatrix dof.MapIndex    // convert axe from degreeGlobal to position in global matrix stiffiners, mass, ...
	degreeForPoint       dof.DoF         // all degree of freedom in global system coordinate for each point
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
func (m *Dim2) AddTrussProperty(beamIndexes ...element.Index) {
	m.truss = append(m.truss, beamIndexes...)
}

// AddSupport - add support for points
func (m *Dim2) AddSupport(support support.Dim2, pointIndexes ...point.Index) {
	for _, p := range pointIndexes {
		m.supports = append(m.supports, supportGroup2d{
			support:    support,
			pointIndex: p,
		})
	}
}

// AddShape - add shape property for beam
func (m *Dim2) AddShape(shape shape.Shape, elements ...element.Index) {
	for _, e := range elements {
		m.shapes = append(m.shapes, shapeGroup{
			shape:        shape,
			elementIndex: e,
		})
	}
}

// AddMaterial - add material for beam
func (m *Dim2) AddMaterial(material material.Linear, elements ...element.Index) {
	for _, e := range elements {
		m.materials = append(m.materials, materialLinearGroup{
			material:     material,
			elementIndex: e,
		})
	}
}

// AddNodeForce - add node force in force case
func (m *Dim2) AddNodeForce(caseNumber int, nodeForce force.NodeDim2, pointIndexes ...point.Index) {
	for i := range m.forceCases {
		if m.forceCases[i].indexCase == caseNumber {
			for _, p := range pointIndexes {
				m.forceCases[i].nodeForces = append(m.forceCases[i].nodeForces, nodeForce2d{
					nodeForce:  nodeForce,
					pointIndex: p,
				})
			}
			return
		}
	}

	for _, p := range pointIndexes {
		nf := nodeForce2d{
			nodeForce:  nodeForce,
			pointIndex: p,
		}
		var fc forceCase2d
		fc.indexCase = caseNumber
		fc.nodeForces = append(fc.nodeForces, nf)
		m.forceCases = append(m.forceCases, fc)
	}

}

/*
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
*/

// GetGlobalDisplacement - return global displacement
func (m *Dim2) GetGlobalDisplacement(caseNumber int, pointIndex point.Index) (d displacement.Dim2, err error) {
	for _, f := range m.forceCases {
		if f.indexCase == caseNumber {
			return f.GetGlobalDisplacement(pointIndex)
		}
	}
	return d, fmt.Errorf("Cannot found case by number")
}

// GetLocalForce - return local force of beam
func (m *Dim2) GetLocalForce(caseNumber int, beamIndex element.Index) (begin, end forceLocal.Dim2, err error) {
	for _, f := range m.forceCases {
		if f.indexCase == caseNumber {
			return f.GetLocalForce(beamIndex)
		}
	}
	return begin, end, fmt.Errorf("Cannot found case by number")
}

// GetReaction - return reaction of support
func (m *Dim2) GetReaction(caseNumber int, pointIndex point.Index) (r reaction.Dim2, err error) {
	for _, f := range m.forceCases {
		if f.indexCase == caseNumber {
			return f.GetReaction(pointIndex)
		}
	}
	return r, fmt.Errorf("Cannot found case by number")
}
