package main

import "testing"

func BenchmarkSidewinder(b *testing.B) {
	sidewinder := MakeSidewinder(NewRand())
	b.ResetTimer()
	for range b.N {
		b.StopTimer()
		grid := NewGrid(100, 100)
		b.StartTimer()

		sidewinder.On(grid)
	}
}
