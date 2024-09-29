package main

import "math/rand"

type BinaryTree struct{}

func (b BinaryTree) On(grid *grid) {
	grid.EachCell(func(c *cell) {
		var neighbors []*cell
		if c.North() != nil {
			neighbors = append(neighbors, c.North())
		}

		if c.East() != nil {
			neighbors = append(neighbors, c.East())
		}

		if len(neighbors) == 0 {
			return
		}

		index := rand.Intn(len(neighbors))
		neighbor := neighbors[index]

		c.Link(neighbor)
	})
}
