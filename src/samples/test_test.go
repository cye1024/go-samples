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

func Benchmark_T67(b *testing.B) {
	for i := 0; i < b.N; i++ {
		T67()
	}
}

func Benchmark_T68(b *testing.B) {
	for i := 0; i < b.N; i++ {
		T68()
	}
}

func Benchmark_T69(b *testing.B) {
	for i := 0; i < b.N; i++ {
		T69()
	}

}

func Benchmark_T70(b *testing.B) {

	for i := 0; i < b.N; i++ {
		T70()
	}
}
