package vector

import (
	"log"
	"math"
	"math/rand"
	"strconv"
	"strings"
)

// VectorDyn implements the dynamic (hence slow) version of the vector interface
type vectorDyn struct {
	cells []float32
}

// Internal 'make' function
func makeDyn(f []float32) Vector {
	var v vectorDyn

	v.cells = make([]float32, len(f))
	for i := range f {
		v.cells[i] = f[i]
	}

	return v
}

// Internal 'zero' function
func zeroDyn(dimension int) Vector {
	var v vectorDyn

	v.cells = make([]float32, dimension)

	return v
}

// Internal 'rand' function
func randDyn(dimension int) Vector {
	var v vectorDyn

	v.cells = make([]float32, dimension)
	for i := range v.cells {
		v.cells[i] = rand.Float32()
	}

	return v
}

// Unit provides a vector with length 1 in the direction of the vector
func (v vectorDyn) Unit() Vector {
	return v.Divs(v.Abs())
}

// Abs provides the euclidian lenth of a vector
// That's how mathematicians specify the absolute value of a vector
func (v vectorDyn) Abs() float32 {

	l := float64(0.0)
	for i := 0; i < v.Len(); i++ {
		l += float64(v.Get(i) * v.Get(i))
	}

	return float32(math.Sqrt(l))
}

// Cbd provides the city-block-distance length of a vector
func (v vectorDyn) Cbd() float32 {
	var l float32

	for i := 0; i < v.Len(); i++ {
		if v.Get(i) > 0 {
			l += v.Get(i)
		} else {
			l -= v.Get(i)
		}
	}

	return l
}

// Add substracts one vector from another
func (v vectorDyn) Add(w Vector) Vector {
	if v.Len() != w.Len() {
		log.Fatalf("Vector.Add: dimensions of vectors must be the same")
	}

	r := Zero(v.Len())
	for i := 0; i < r.Len(); i++ {
		r = r.Set(i, v.Get(i)+w.Get(i))
	}

	return r
}

// Sub substracts one vector from another
func (v vectorDyn) Sub(w Vector) Vector {
	if v.Len() != w.Len() {
		log.Fatalf("Vector.Sub: dimensions of vectors must be the same")
	}

	r := Zero(v.Len())
	for i := 0; i < r.Len(); i++ {
		r = r.Set(i, v.Get(i)-w.Get(i))
	}

	return r
}

// Min provides a vector containing the smallest elements of both vectors
func (v vectorDyn) Min(w Vector) Vector {
	if v.Len() != w.Len() {
		log.Fatalf("Vector.Min: dimensions of vectors must be the same")
	}

	r := Zero(v.Len())
	for i := 0; i < r.Len(); i++ {
		if v.Get(i) < w.Get(i) {
			r = r.Set(i, v.Get(i))
		} else {
			r = r.Set(i, w.Get(i))
		}
	}

	return r
}

// MinD provides the index of the smallest value in t he vector
func (v vectorDyn) MinD() int {

	r := 0
	for i := 0; i < v.Len(); i++ {
		if v.Get(i) < v.Get(r) {
			r = i
		}
	}

	return r
}

// Max set every element of the resulting vector to the highest option
// It provides the first occurence if there are more dimensions with this value
func (v vectorDyn) Max(w Vector) Vector {
	if v.Len() != w.Len() {
		log.Fatalf("Vector.Max: dimensions of vectors must be the same")
	}

	r := Zero(v.Len())
	for i := 0; i < r.Len(); i++ {
		if v.Get(i) > w.Get(i) {
			r = r.Set(i, v.Get(i))
		} else {
			r = r.Set(i, w.Get(i))
		}
	}

	return r
}

// MaxD provides the index of the largest value in the vector
// It provides the first occurence if there are more dimensions with this value
func (v vectorDyn) MaxD() int {

	r := 0
	for i := 0; i < v.Len(); i++ {
		if v.Get(i) > v.Get(r) {
			r = i
		}
	}

	return r
}

// Muls multiplies a vector by a scalar.
func (v vectorDyn) Muls(s float32) Vector {

	r := Zero(v.Len())
	for i := 0; i < r.Len(); i++ {
		r = r.Set(i, v.Get(i)*s)
	}

	return r
}

// Divs divides a vector by a scalar.
func (v vectorDyn) Divs(s float32) Vector {

	r := Zero(v.Len())
	for i := 0; i < r.Len(); i++ {
		r = r.Set(i, v.Get(i)/s)
	}

	return r
}

// Len retrieves the number of dimensions
func (v vectorDyn) Len() int {
	return len(v.cells)
}

// Raw retrieves the values of the vector as a slice
func (v vectorDyn) Raw() []float32 {
	return v.cells
}

// Get retrieves the value of a single cell
func (v vectorDyn) Get(i int) float32 {
	if (i < 0) || (i >= len(v.cells)) {
		log.Fatalf("Vector.Get: Index out of bounds")
	}

	return v.cells[i]
}

// Set changes the value of a single cell
func (v vectorDyn) Set(i int, f float32) Vector {
	if (i < 0) || (i >= len(v.cells)) {
		log.Fatalf("Vector.Set: Index out of bounds")
	}

	v.cells[i] = f
	return v
}

// String() implements the Stringer interface
func (v vectorDyn) String() string {
	var s strings.Builder

	s.WriteString("[")
	for i, f := range v.cells {
		if i > 0 {
			s.WriteString(", ")
		}
		s.WriteString(strconv.FormatFloat(float64(f), 'f', 3, 64))
	}
	s.WriteString("]")

	return s.String()
}
