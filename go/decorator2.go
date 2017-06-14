package main

import (
	"fmt"
	"reflect"
	"runtime"
	"time"
)

type SumFunc func(start, end int64) int64

func getFuncName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func timedSumFunc(f SumFunc) SumFunc {
	return func(start, end int64) int64 {
		defer func(t time.Time) {
			fmt.Printf("---Time Elapsed (%s): %v ---\n", getFuncName(f), time.Since(t))
		}(time.Now())

		return f(start, end)
	}
}

func Sum1(start, end int64) int64 {
	var sum int64
	for i := start; i <= end; i++ {
		sum += i
	}

	return sum
}

func Sum2(start, end int64) int64 {
	// mathmatic solution

	return (end - start + 1) * (start + end) / 2
}

func main() {
	sum1 := timedSumFunc(Sum1)
	sum2 := timedSumFunc(Sum2)

	fmt.Printf("%d %d\n", sum1(-100000, 10000000), sum2(-100000, 10000000))
}
