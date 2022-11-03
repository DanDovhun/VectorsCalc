package two_dim

import (
	"fmt"
	"math"

	"github.com/DanDovhun/VectorsCalc/angle"
)

type Vector struct {
	Name string
	X    float64
	Y    float64
}

// New serves as a constructor for the Vector struct
func New(name string, x, y float64) Vector {
	return Vector{
		Name: name,
		X:    x,
		Y:    y,
	}
}

// Magnitude of a vector
func (vec Vector) Magnitude() float64 {
	return math.Sqrt(math.Pow(vec.X, 2) + math.Pow(vec.Y, 2))
}

// Elevation: angle of elevation of the angle
func (vec Vector) Elevation() angle.Angle {
	return angle.New(math.Atan(vec.Y / vec.X))
}

// AngleBetween vecotrs
func (vec Vector) AngleBetween(vecB Vector) angle.Angle {
	prod := vec.DotProduct(vecB)
	magA := vec.Magnitude()
	magB := vecB.Magnitude()

	return angle.New(math.Acos(prod / (magA * magB)))
}

// Add a vector
func (vec Vector) Add(vecB Vector) Vector {
	x := vec.X + vecB.X
	y := vec.Y + vecB.Y

	return New("add", x, y)
}

// Sub: Subtract vectors
func (vec Vector) Sub(vecB Vector) Vector {
	x := vec.X - vecB.X
	y := vec.Y - vecB.Y

	return New("Sub", x, y)
}

func (vec Vector) DotProduct(vecB Vector) float64 {
	return (vec.X * vecB.X) + (vec.Y * vecB.Y)
}

func (vec Vector) CrossProduct(vecB Vector) float64 {
	return (vec.X * vecB.Y) - (vec.Y * vecB.X)
}

func (vec Vector) ToString() string {
	return fmt.Sprintf("%v(%v, %v)", vec.Name, vec.X, vec.Y)
}
