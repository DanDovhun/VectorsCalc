package three_dim

import (
	"fmt"
	"math"

	"github.com/DanDovhun/VectorsCalc/angle"
)

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

func (vec Vector) GetProperties() Properties {
	props := NewProperties(vec)

	return props
}

func (vec Vector) Add(vecB Vector) Vector {
	x := vec.X + vecB.X
	y := vec.Y + vecB.Y
	z := vec.Z + vecB.Z

	return New("Add", x, y, z)
}

func (vec Vector) Sub(vecB Vector) Vector {
	x := vec.X - vecB.X
	y := vec.Y - vecB.Y
	z := vec.Z - vecB.Z

	return New("Sub", x, y, z)
}

func (vec Vector) CrossProduct(vecB Vector) Vector {
	x := (vec.Y * vecB.Z) - (vec.Z * vecB.Y)
	y := (vec.Z * vecB.X) - (vec.X * vecB.Z)
	z := (vec.X * vecB.Y) - (vec.Y * vecB.X)

	return New("Dot Prod", x, y, z)
}

func (vec Vector) DotProduct(vecB Vector) float64 {
	return (vec.X * vecB.X) + (vec.Y*vecB.Y)*(vec.Z*vecB.Z)
}

func (vec Vector) AngleBetween(vecB Vector) angle.Angle {
	return angle.New(math.Acos(vec.DotProduct(vecB) / (vec.Magnitude() * vecB.Magnitude())))
}

func (vec Vector) ToString() string {
	return fmt.Sprintf("%v(%v, %v, %v)", vec.Name, vec.X, vec.Y, vec.Z)
}
