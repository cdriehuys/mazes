package main

import (
	"os"

	"github.com/cdriehuys/mazes"
)

func main() {
	rand := mazes.NewRand()

	grid := mazes.NewGrid(4, 4)
	mazes.MakeSidewinder(rand).On(grid)

	mazes.MakeASCIIRenderer(os.Stdout).Render(grid)
}
