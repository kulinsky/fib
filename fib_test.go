package fib_test

import (
	"testing"

	"github.com/kulinsky/fib"
)

type fibTest struct {
	arg, expected int
}

var fibTests = []fibTest{
	{0, 0},
	{1, 1},
	{2, 1},
	{3, 2},
	{10, 55},
	{20, 6765},
}

func TestFibonacci(t *testing.T) {
	for _, test := range fibTests {
		if output := fib.Fibonacci(test.arg); output != test.expected {
			t.Errorf("Output %d not equal to expected %d", output, test.expected)
		}
	}
}

func TestTabulation(t *testing.T) {
	for _, test := range fibTests {
		if output := fib.Tabulation(test.arg); output != test.expected {
			t.Errorf("Output %d not equal to expected %d", output, test.expected)
		}
	}
}

func TestMemoization(t *testing.T) {
	for _, test := range fibTests {
		if output := fib.Memoization(test.arg); output != test.expected {
			t.Errorf("Output %d not equal to expected %d", output, test.expected)
		}
	}
}

func BenchmarkTabulation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib.Tabulation(20)
	}
}

func BenchmarkMemoization(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib.Memoization(20)
	}
}
