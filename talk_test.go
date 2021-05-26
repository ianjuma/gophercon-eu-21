package gophercon

import "testing"

// we can verify this by running a benchmark
// remove fmt.Println - escapes
var list []int

// passing down - manipulating a pointer does NOT escape
func BenchmarkPassDownRef(b *testing.B) {
	for i := 0; i < b.N; i++ {
	}
}

// passing up - return a pointer escapes
func BenchmarkPassUpRef(b *testing.B) {
	for i := 0; i < b.N; i++ {
	}
}

// interface example
// passing up example of an escape - returns
// allocation on the heap - escape
func BenchmarkPull1(b *testing.B) {
	for i := 0; i < b.N; i++ {
	}
}

func BenchmarkPull2(b *testing.B) {
	_ = make([]int, 1024)
	for i := 0; i < b.N; i++ {
	}
}

func BenchmarkLiteralFunctions(b *testing.B) {
	for i := 0; i < b.N; i++ {
	}
}
