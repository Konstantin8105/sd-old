package dof

import (
	"fmt"
	"sort"

	"github.com/Konstantin8105/GoFea/utils"
)

type row struct {
	axe      int
	position int
}

// MapIndex - map of indexes
type MapIndex struct {
	rows []row
}

// NewMapIndex - constructor for index map
func NewMapIndex(axes *[]AxeNumber) (m MapIndex) {
	a := ConvertToInt(*axes)
	utils.UniqueInt(&a)
	size := len(a)
	m.rows = make([]row, size, size)
	for i := 0; i < size; i++ {
		m.rows[i] = row{
			axe:      a[i],
			position: i,
		}
	}
	return m
}

// GetByAxe - return index by axe
func (m *MapIndex) GetByAxe(axe AxeNumber) (int, error) {
	a := int(axe)
	i := sort.Search(len(m.rows), func(i int) bool { return m.rows[i].axe >= a })
	if i >= 0 && i < len(m.rows) && m.rows[i].axe == a {
		// index is present at array[i]
		return m.rows[i].position, nil
	}
	// index is not present in array,
	// but i is the index where it would be inserted.
	return -1, fmt.Errorf("Not found")
}
