package main

import (
	"fmt"
	"runtime"
	"reflect"
	"time"
)

type SumFunc func(start, end int64) int64

func Sum1(start, end int64) int64 {
	return (end - start + 1) * (start + end) / 2
}

func Sum2(start, end int64) int64 {
	var sum int64

	for i := start; i <=end; i++ {
		sum += i
	}

	return sum
}

func DecoratSumFunc(f SumFunc) SumFunc {
	return func(start, end int64) int64 {
		defer func(startTime time.Time) {
			fmt.Printf("func name (%s) elapse %v\n", getFuncName(f), time.Since(startTime))
		}(time.Now())

		return f(start, end)
	}
}

func getFuncName(f interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}

func main() {
	fmt.Printf("%d ---- %d\n", DecoratSumFunc(Sum1)(-10, 100000), DecoratSumFunc(Sum2)(-10, 100000))
}
