package game

// Grid is a grid of values of any type.
type Grid map[Position]CellType

// NewGrid will create a new empty grid.
func NewGrid() Grid {
	return Grid{}
}

// Add value at position (x, y).
func (g Grid) Add(x, y int, value CellType) {
	g[NewPosition(x, y)] = value
}

// Retrieve value at position (x, y).
func (g Grid) Retrieve(x, y int) (interface{}, bool) {
	value, ok := g[NewPosition(x, y)]
	return value, ok
}

// Delete value at position (x, y).
func (g Grid) Delete(x, y int) {
	delete(g, NewPosition(x, y))
}

// Clear the entire grid.
func (g Grid) Clear() {
	for k := range g {
		delete(g, k)
	}
}

// Copy the grid structure using a shallow copy.
func (g Grid) Copy() Grid {
	copy := NewGrid()

	for k, v := range g {
		copy[k] = v
	}

	return copy
}
