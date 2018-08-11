package main

import (
	"testing"
	"fmt"
)

func TestCompute(t *testing.T) {
	rectCount := 8
	fmt.Println(compute(rectCount))
	fmt.Println(computeC(rectCount, 6))
}
