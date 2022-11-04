package angle

import "math"

type Angle struct {
	Radians float64
	Degrees float64
}

var New = func(radians float64) Angle {
	degrees := radians * 180 / math.Pi

	return Angle{
		Radians: radians,
		Degrees: degrees,
	}
}
