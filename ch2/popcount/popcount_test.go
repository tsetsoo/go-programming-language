package popcount

import (
	"testing"
)

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(42947295)
	}
}

func BenchmarkPopCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountLoop(42947295)
	}
}

func BenchmarkPopCountShift64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountShift64(42947295)
	}
}

func BenchmarkPopCountByShifting(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByShifting(42947295)
	}
}

func BenchmarkPopCountByClear(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByClear(42947295)
	}
}
