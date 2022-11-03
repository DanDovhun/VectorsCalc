package three_dim

import "math"

type Vector struct {
	Name string
	X    float64
	Y    float64
	Z    float64
}

func New(name string, x, y, z float64) Vector {
	return Vector{
		Name: name,
		X:    x,
		Y:    y,
		Z:    z,
	}
}

func (vec Vector) Magnitude() float64 {
	return math.Sqrt(math.Pow(vec.X, 2) + math.Pow(vec.Y, 2) + math.Pow(vec.Z, 2))
}

func (vec Vector) GetProperties()
