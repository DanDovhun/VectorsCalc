package three_dim

type Vector struct {
	Name string
	X    float64
	Y    float64
	Z    float64
}

func New(name string, x, y, z float64) {
	return Vector
}
