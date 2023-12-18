package main

import (
	"fmt"
	"math"
)

func main() {
	var a, v0, s0, t float64
	fmt.Print("Enter acceleration: ")
	fmt.Scan(&a)
	fmt.Print("Enter initial velocity: ")
	fmt.Scan(&v0)
	fmt.Print("Enter initial displacement: ")
	fmt.Scan(&s0)
	fmt.Print("Enter time: ")
	fmt.Scan(&t)
	DisplaceAtTime := GenDisplaceFn(a, v0, s0)
	fmt.Printf("Displacement at %.3f s is %.3f m.\n", t, DisplaceAtTime(t))
}

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	return func(t float64) float64 {
		return 0.5*a*math.Pow(t, 2) + v0*t + s0
	}
}
