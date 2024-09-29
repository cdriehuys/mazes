package main

import (
	"fmt"
	"io"
)

// asciiRenderer outputs the grid as ASCII symbols to the provided output
// writer.
type asciiRenderer struct {
	output io.Writer
}

func (r asciiRenderer) Render(grid *grid) {
	// Top boundary
	r.fprint("+")
	for i := 0; i < grid.columns; i++ {
		r.fprint("---+")
	}
	r.fprint("\n")

	for rowNum := 0; rowNum < grid.rows; rowNum++ {
		bodyLine := "|"
		borderLine := "+"

		for colNum := 0; colNum < grid.columns; colNum++ {
			cell := grid.Cell(rowNum, colNum)

			bodyLine += "   "
			if cell.Linked(cell.East()) {
				bodyLine += " "
			} else {
				bodyLine += "|"
			}

			if cell.Linked(cell.South()) {
				borderLine += "   "
			} else {
				borderLine += "---"
			}
			borderLine += "+"
		}

		r.fprintln(bodyLine)
		r.fprintln(borderLine)
	}
}

func (r asciiRenderer) fprint(a ...any) {
	fmt.Fprint(r.output, a...)
}

func (r asciiRenderer) fprintln(a ...any) {
	fmt.Fprintln(r.output, a...)
}
