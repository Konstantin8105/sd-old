package dof

import (
	"fmt"
	"sort"
)

type row struct {
	axe      AxeNumber // int
	position int
}

// MapIndex - map of indexes
type MapIndex struct {
	rows []row
}

// NewMapIndex - constructor for index map
func NewMapIndex(a *[]AxeNumber) (m MapIndex) {
	size := len(*a)
	m.rows = make([]row, size, size)
	for i := 0; i < size; i++ {
		m.rows[i] = row{
			axe:      (*a)[i],
			position: i,
		}
	}
	return m
}

// GetByAxe - return index by axe
func (m *MapIndex) GetByAxe(axe AxeNumber) (int, error) {
	a := int(axe)
	i := sort.Search(len(m.rows), func(i int) bool { return int(m.rows[i].axe) >= a })
	if i >= 0 && i < len(m.rows) && int(m.rows[i].axe) == a {
		// index is present at array[i]
		return m.rows[i].position, nil
	}
	// index is not present in array,
	// but i is the index where it would be inserted.
	return -1, fmt.Errorf("Not found")
}
