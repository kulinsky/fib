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

func TestFib(t *testing.T) {
	for _, test := range fibTests {
		if output := fib.Fib(test.arg); output != test.expected {
			t.Errorf("Output %d not equal to expected %d", output, test.expected)
		}
	}
}
