package main

import (
	"os"
)

type cell struct {
	row    int
	column int
	links  map[*cell]bool

	north *cell
	east  *cell
	south *cell
	west  *cell
}

func newCell(row, column int) *cell {
	return &cell{
		row:    row,
		column: column,
		links:  make(map[*cell]bool),
	}
}

func (c *cell) Link(cell *cell) {
	c.links[cell] = true
	cell.links[c] = true
}

func (c *cell) Unlink(cell *cell) {
	delete(c.links, cell)
	delete(cell.links, c)
}

func (c *cell) Links() []*cell {
	links := make([]*cell, len(c.links))
	for c := range c.links {
		links = append(links, c)
	}

	return links
}

func (c *cell) Linked(other *cell) bool {
	if other == nil {
		return false
	}

	_, exists := c.links[other]

	return exists
}

func (c *cell) North() *cell {
	return c.north
}

func (c *cell) East() *cell {
	return c.east
}

func (c *cell) South() *cell {
	return c.south
}

func (c *cell) West() *cell {
	return c.west
}

func (c *cell) Neighbors() []*cell {
	neighbors := []*cell{}

	if c.north != nil {
		neighbors = append(neighbors, c.north)
	}

	if c.east != nil {
		neighbors = append(neighbors, c.east)
	}

	if c.south != nil {
		neighbors = append(neighbors, c.south)
	}

	if c.west != nil {
		neighbors = append(neighbors, c.west)
	}

	return neighbors
}

type grid struct {
	rows    int
	columns int

	cells [][]*cell
}

func NewGrid(rows, columns int) *grid {
	gridRows := make([][]*cell, rows)
	for row := 0; row < rows; row++ {
		gridCols := make([]*cell, columns)
		for col := 0; col < columns; col++ {
			gridCols[col] = newCell(row, col)
		}

		gridRows[row] = gridCols
	}

	// Do a second pass to set up all the links
	for row := 0; row < rows; row++ {
		for col := 0; col < columns; col++ {
			cell := gridRows[row][col]

			if row-1 >= 0 {
				cell.north = gridRows[row-1][col]
			}

			if col+1 < columns {
				cell.east = gridRows[row][col+1]
			}

			if row+1 < rows {
				cell.south = gridRows[row+1][col]
			}

			if col-1 >= 0 {
				cell.west = gridRows[row][col-1]
			}
		}
	}

	return &grid{
		rows:    rows,
		columns: columns,

		cells: gridRows,
	}
}

func (g *grid) Cell(row, col int) *cell {
	return g.cells[row][col]
}

func (g *grid) EachCell(visit func(c *cell)) {
	for row := 0; row < g.rows; row++ {
		for col := 0; col < g.columns; col++ {
			visit(g.cells[row][col])
		}
	}
}

func (g *grid) EachRow(visit func(row []*cell)) {
	for row := range g.rows {
		visit(g.cells[row])
	}
}

func main() {
	rand := NewRand()

	g := NewGrid(10, 10)
	MakeSidewinder(rand).On(g)

	asciiRenderer{os.Stdout}.Render(g)
}
