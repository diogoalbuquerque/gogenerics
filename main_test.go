package main

import (
	"fmt"
	"testing"
)

func Benchmark_sameEqualsInt(t *testing.B) {
	_ = same(1, 1)
}

func Benchmark_sameNotEqualsInt(t *testing.B) {
	_ = same(1, 2)
}
func Benchmark_sameEqualsFloat(t *testing.B) {
	_ = same(1.99, 1.99)
}

func Benchmark_sameNotEqualsFloat(t *testing.B) {
	_ = same(1.99, 2.99)
}

func Benchmark_miniAndMaxi(t *testing.B) {
	_, _ = miniAndMaxi([]int{1, 5, 3, 2, 4})
}

func Benchmark_onlyOneString(t *testing.B) {
	_ = onlyOne([]string{"Banana", "Carrot", "Guava", "Pineapple", "Carrot", "Guava"})
}

func Benchmark_onlyOneFloat(t *testing.B) {
	_ = onlyOne([]float64{1.12, 2.21, 1.12, 1.99, 9.87, 9.87, 12.12})
}

func Fuzz_same(f *testing.F) {
	f.Add(1, 2)
	f.Add(1, 1)
	f.Add(2, 1)
	f.Fuzz(func(t *testing.T, data1 int, data2 int) {
		fmt.Println(data1)
		fmt.Println(data2)
		same(data1, data2)
	})
}
