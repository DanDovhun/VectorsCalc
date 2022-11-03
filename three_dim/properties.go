package three_dim

import (
	"math"

	"github.com/DanDovhun/VectorsCalc/angle"
)

type Properties struct {
	Alpha     angle.Angle
	Beta      angle.Angle
	Gamma     angle.Angle
	Elevation angle.Angle
	Rotation  angle.Angle
}

func NewProperties(vec Vector) Properties {
	var props Properties

	props.Alpha = setAlpha(vec)
	props.Beta = setBeta(vec)
	props.Gamma = setGamma(vec)
	props.Elevation = setElevation(vec)
	props.Rotation = setRotation(vec)

	return props
}

func setAlpha(vec Vector) angle.Angle {
	u := vec.Magnitude()

	return angle.New(math.Acos(vec.X / u))
}

func setBeta(vec Vector) angle.Angle {
	u := vec.Magnitude()
	return angle.New(math.Acos(vec.Y / u))
}

func setGamma(vec Vector) angle.Angle {
	u := vec.Magnitude()
	return angle.New(math.Acos(vec.Z / u))
}

func setRotation(vec Vector) angle.Angle {
	v := math.Sqrt(math.Pow(vec.X, 2) + math.Pow(vec.Y, 2))
	return angle.New(math.Acos(vec.X / v))
}

func setElevation(vec Vector) angle.Angle {
	u := vec.Magnitude()
	v := math.Sqrt(math.Pow(vec.X, 2) + math.Pow(vec.Y, 2))

	return angle.New(math.Acos(v / u))
}

func (props Properties) GetProperties() (angle.Angle, angle.Angle, angle.Angle, angle.Angle, angle.Angle) {
	return props.Alpha, props.Beta, props.Gamma, props.Elevation, props.Rotation
}
