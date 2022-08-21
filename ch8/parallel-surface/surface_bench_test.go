package main

import "testing"

// BenchmarkSurface-16                         6102            199806 ns/op
// BenchmarkParallelSurface-16                31147             40679 ns/op
// BenchmarkParallelJSurface-16                 344           3476127 ns/op

func BenchmarkSurface(b *testing.B) {
	for i := 0; i < b.N; i++ {
		surface()
	}
}

func BenchmarkParallelSurface(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parallelSurface()
	}
}

func BenchmarkParallelJSurface(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parallelJSurface()
	}
}
