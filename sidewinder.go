package main

type Sidewinder struct {
	rand random
}

func MakeSidewinder(rand random) Sidewinder {
	return Sidewinder{rand}
}

func (s Sidewinder) On(grid *grid) {
	grid.EachRow(func(row []*cell) {
		var run []*cell

		for _, cell := range row {
			run = append(run, cell)

			at_east_bound := cell.East() == nil
			at_north_bound := cell.North() == nil

			should_close := at_east_bound ||
				(!at_north_bound && s.rand.Bool())

			if should_close {
				member := s.rand.Cell(run)
				if member.North() != nil {
					member.Link(member.North())
				}
				run = nil
			} else {
				cell.Link(cell.East())
			}
		}
	})
}
