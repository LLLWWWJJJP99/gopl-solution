// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package popcount_test

import (
	"awesomeProject/execise/exe9_2"
	"testing"
)

// -- Benchmarks --
func bench(b *testing.B, f func(uint64) int) {
	for i := 0; i < b.N; i++ {
		f(uint64(i))
	}
}

func BenchmarkPopCount(b *testing.B) {
	bench(b, popcount.PopCount)
}

// Go 1.6, 2.67GHz Xeon
// $ go test -cpu=4 -bench=. gopl.io/ch2/popcount
// BenchmarkPopCount-4                  200000000         6.30 ns/op
// BenchmarkBitCount-4                  300000000         4.15 ns/op
// BenchmarkPopCountByClearing-4        30000000         45.2 ns/op
// BenchmarkPopCountByShifting-4        10000000        153 ns/op
//
// Go 1.6, 2.5GHz Intel Core i5
// $ go test -cpu=4 -bench=. gopl.io/ch2/popcount
// BenchmarkPopCount-4                  200000000         7.52 ns/op
// BenchmarkBitCount-4                  500000000         3.36 ns/op
// BenchmarkPopCountByClearing-4        50000000         34.3 ns/op
// BenchmarkPopCountByShifting-4        20000000        108 ns/op
//
// Go 1.7, 3.5GHz Xeon
// $ go test -cpu=4 -bench=. gopl.io/ch2/popcount
// BenchmarkPopCount-12                 2000000000        0.28 ns/op
// BenchmarkBitCount-12                 2000000000        0.27 ns/op
// BenchmarkPopCountByClearing-12       100000000        18.5 ns/op
// BenchmarkPopCountByShifting-12       20000000         70.1 ns/op