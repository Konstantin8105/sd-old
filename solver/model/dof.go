package model

import (
	"github.com/Konstantin8105/GoFea/input/point"
	"github.com/Konstantin8105/GoFea/solver/dof"
	"github.com/Konstantin8105/GoFea/solver/finiteElement"
)

// generateDof - create degree's of freedom for model
func (m *Dim2) generateDof() error {

	// Generate degreeForPoint - degree of freedom in global system for each point of finite element
	{
		pointIndexes := make([]point.Index, len(m.points))
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
			fe, err := m.getFiniteElement(ele.GetIndex())
			if err != nil {
				return err
			}
			_, localAxes := finiteElement.GetStiffinerGlobalK(fe, &m.degreeForPoint, finiteElement.WithoutZeroStiffiner)
			axes = append(axes, localAxes...)
		}
		dof.UniqueAxeNumber(&axes)
		m.degreeInGlobalMatrix = axes
	}

	// Create convertor axe number to position in global matrix
	m.indexsInGlobalMatrix = dof.NewMapIndex(&m.degreeInGlobalMatrix)

	return nil
}
