package prime

import (
	"testing"
)

func BenchmarkFindPrimesBasic100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FindPrimesBasic(1000000000)
	}
}

func BenchmarkFindPrimesOptimized100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FindPrimesOptimized(1000000000)
	}
}

func TestFindPrimesOptimized(t *testing.T) {
	FindPrimesOptimized(30)
}
