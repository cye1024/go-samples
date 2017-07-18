package samples

import "testing"

func Benchmark_T40(b *testing.B) {
	for i := 0; i < b.N; i++ {
		T40()
	}
}

func Benchmark_T41(b *testing.B) {
	for i := 0; i < b.N; i++ {
		T41()
	}
}
