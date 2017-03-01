package echo

import "testing"

func BenchmarkEcho2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Echo2()
	}
}

func BenchmarkEcho3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Echo3()
	}
}
