package main

import (
	"fmt"
	"math/cmplx"
	"math"
)

var (
	KnowTorF bool = true;
	MaxInt uint64 = 1<<64 -1;
	z complex128 = cmplx.Sqrt(-5 + 12i);
)

func main() {
	fmt.Println(math.Pi)
	fmt.Printf("type: %T value: %v\n", KnowTorF, KnowTorF)
	fmt.Printf("type: %T value: %v\n", MaxInt, MaxInt)
	fmt.Printf("type: %T value: %v\n", z, z)
}
