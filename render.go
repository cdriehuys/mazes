package mazes

import (
	"fmt"
	"io"
)

// ASCIIRenderer outputs the grid as ASCII symbols to the provided output
// writer.
type ASCIIRenderer struct {
	output io.Writer
}

func MakeASCIIRenderer(output io.Writer) ASCIIRenderer {
	return ASCIIRenderer{output}
}

func (r ASCIIRenderer) Render(grid *Grid) {
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

func (r ASCIIRenderer) fprint(a ...any) {
	fmt.Fprint(r.output, a...)
}

func (r ASCIIRenderer) fprintln(a ...any) {
	fmt.Fprintln(r.output, a...)
}
