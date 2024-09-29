package main

import "testing"

func BenchmarkBinaryTree(b *testing.B) {
	binaryTree := MakeBinaryTree(NewRand())
	b.ResetTimer()
	for range b.N {
		b.StopTimer()
		grid := NewGrid(100, 100)
		b.StartTimer()

		binaryTree.On(grid)
	}
}
