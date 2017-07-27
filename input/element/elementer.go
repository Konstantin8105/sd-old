package element

var _ Elementer = (*Beam)(nil)

// Elementer - interface typical for elements
type Elementer interface {
	// GetIndex - return index of element
	GetIndex() ElementIndex
}
