package angle

import "math"

type Angle struct {
	Radians float64
	Degrees float64
}

func New(radians float64) Angle {
	degrees := radians * 180 / math.Pi

	return Angle{
		Radians: radians,
		Degrees: degrees,
	}
}
