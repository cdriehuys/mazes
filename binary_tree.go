package main

type random interface {
	Bool() bool
	Cell([]*cell) *cell
}

type BinaryTree struct {
	rand random
}

func MakeBinaryTree(rand random) BinaryTree {
	return BinaryTree{rand}
}

func (b BinaryTree) On(grid *grid) {
	grid.EachCell(func(c *cell) {
		var neighbors []*cell
		if c.North() != nil {
			neighbors = append(neighbors, c.North())
		}

		if c.East() != nil {
			neighbors = append(neighbors, c.East())
		}

		if neighbor := b.rand.Cell(neighbors); neighbor != nil {
			c.Link(neighbor)
		}
	})
}
