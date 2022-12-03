package math

import "math"

type Vector [2]float64

func (v *Vector) Normalize() *Vector {
	strength := v.Length()
	v[0] = v[0] / strength
	v[1] = v[1] / strength

	return v
}

func (v *Vector) Length() float64 {
	return math.Sqrt(v[0]*v[0] + v[1]*v[1])
}
