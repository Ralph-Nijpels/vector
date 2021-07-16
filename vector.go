package vector

import (
	"fmt"
)

// Vector implements a simple mathematical vector
type Vector interface {
	Unit() Vector
	Abs() float32
	Cbd() float32
	Add(w Vector) Vector
	Sub(w Vector) Vector
	Min(w Vector) Vector
	MinD() int
	Max(w Vector) Vector
	MaxD() int
	Muls(s float32) Vector
	Divs(s float32) Vector
	Len() int
	Raw() []float32
	Get(i int) float32
	Set(i int, f float32) Vector
	fmt.Stringer
}

// Make creates a vector based on a list of values
func Make(f []float32) Vector {
	if len(f) == 3 {
		return make3D(f)
	}
	return makeDyn(f)
}

// Zero creates a vector of the requested size set to the origin
func Zero(dimension int) Vector {
	if dimension == 3 {
		return zero3D()
	}
	return zeroDyn(dimension)
}

// Rand creates a vector of the requested size set to a random location
func Rand(dimension int) Vector {
	if dimension == 3 {
		return rand3D()
	}
	return randDyn(dimension)
}
