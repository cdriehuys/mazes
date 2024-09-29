package mazes

import (
	"math/rand"
	"time"
)

type Rand struct {
	r *rand.Rand
}

func NewRand() *Rand {
	return &Rand{
		r: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (r Rand) Bool() bool {
	return r.r.Intn(2) == 0
}

func (r Rand) Cell(cells []*cell) *cell {
	if len(cells) == 0 {
		return nil
	}

	return cells[r.r.Intn(len(cells))]
}
