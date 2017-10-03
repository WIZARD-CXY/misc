package main

import (
	"fmt"
	"math/bits"
	"testing"
)

func NaiveOnesCount64(x uint64) int {
	var c, i uint64

	for i = 0; i < 64; i++ {
		c += (x >> i) & 1
	}

	return int(c)
}

var numbers = []uint64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
	100, 200, 300, 400, 500, 600, 700, 800, 900,
	11111, 22222, 33333, 44444, 55555, 66666, 77777, 88888, 99999,
	190239012390, 123901312, 6549654, 45904059,
}

var Output int

func TestNaiveOnesCount(t *testing.T) {
	m := len(numbers)

	for n := 0; n < m; n++ {
		x := numbers[n]

		got := NaiveOnesCount64(x)
		want := bits.OnesCount64(x)

		if got != want {
			t.Errorf("naive bit count does not right, want %d, got %d", want, got)
		}
	}
}

func BenchmarkNaiveOnesCount64(b *testing.B) {
	fmt.Println("xixi", b.N)

	for n := 0; n < b.N; n++ {
		Output += NaiveOnesCount64(numbers[n&31])
	}
}

func BenchmarkBitsOnesCount64(b *testing.B) {
	fmt.Println("haha", b.N)
	for n := 0; n < b.N; n++ {

		Output += bits.OnesCount64(numbers[n&31])
	}
}
