package model

import (
	"github.com/Konstantin8105/GoFea/input/element"
	"github.com/Konstantin8105/GoFea/input/point"
	"github.com/Konstantin8105/GoFea/solver/dof"
	"github.com/Konstantin8105/GoFea/solver/finiteElement"
	"github.com/Konstantin8105/GoFea/utils"
)

// generateDof - create degree's of freedom for model
func (m *Dim2) generateDof() {

	// Generate degreeForPoint - degree of freedom in global system for each point of finite element
	{
		pointIndexes := make([]point.Index, len(m.points), len(m.points))
		for i := range m.points {
			pointIndexes = append(pointIndexes, m.points[i].Index)
		}
		m.degreeForPoint = dof.DoF{
			DofArray:  pointIndexes,
			Dimension: dof.Dim2d,
		}
	}

	// Generate degreeInGlobalMatrix - degree of freedom for creating global matrix of stiffnes, mass, ...
	{
		var axes []dof.AxeNumber
		for _, ele := range m.elements {
			switch ele.(type) {
			case element.Beam:
				beam := ele.(element.Beam)
				fe := m.getBeamFiniteElement(beam.Index)
				_, localAxes := finiteElement.GetStiffinerGlobalK(fe, &m.degreeForPoint, finiteElement.WithoutZeroStiffiner)
				axes = append(axes, localAxes...)
			default:
				panic("")
			}
		}
		utils.UniqueAxeNumber(&axes)
		m.degreeInGlobalMatrix = axes
	}

	// Create convertor axe number to position in global matrix
	m.indexsInGlobalMatrix = dof.NewMapIndex(&m.degreeInGlobalMatrix)
	return
}