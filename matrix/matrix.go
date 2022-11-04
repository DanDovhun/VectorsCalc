package matrix

import (
	"errors"
	"fmt"
)

type Matrix struct {
	Height int
	Width  int
	Matrix [][]float64
}

var NewBlank = func(height, width int) Matrix {
	return Matrix{
		Height: height,
		Width:  width,
	}
}

var New = func(arr [][]float64) (Matrix, error) {
	width := len(arr[0])

	var isValid = func(arr [][]float64) bool {
		for _, i := range arr {
			if len(i) != width {
				return false
			}
		}

		return true
	}

	if !isValid(arr) {
		return Matrix{}, errors.New("All rows must be the same length")
	}

	return Matrix{
		Height: width,
		Width:  len(arr),
		Matrix: arr,
	}, nil
}

var canBeAdded = func(a, b Matrix) bool {
	if a.Height == b.Height && a.Width == b.Width {
		return true
	}

	return false
}

var canBeMultiplied = func(a, b Matrix) bool {
	if a.Height == b.Width && a.Width == b.Height {
		return true
	}

	return false
}

func (a Matrix) Add(b Matrix) (Matrix, error) {
	if !canBeAdded(a, b) {
		return Matrix{}, errors.New("Both matrixes must have equal dimensions")
	}

	var arr [][]float64

	for i := 0; i < a.Height; i++ {
		var row []float64

		for j := 0; j < a.Width; j++ {
			row = append(row, a.Matrix[i][j]+b.Matrix[i][j])
		}

		arr = append(arr, row)
	}

	return New(arr)
}

func (a Matrix) Sub(b Matrix) (Matrix, error) {
	if !canBeAdded(a, b) {
		return Matrix{}, errors.New("Both matrixes must have equal dimensions")
	}

	var arr [][]float64

	for i := 0; i < a.Height; i++ {
		var row []float64

		for j := 0; j < a.Width; j++ {
			row = append(row, a.Matrix[i][j]-b.Matrix[i][j])
		}

		arr = append(arr, row)
	}

	return New(arr)
}

func (a Matrix) Mul(b Matrix) (Matrix, error) {
	if !canBeMultiplied(a, b) {
		return Matrix{}, errors.New("Matrix A's width must be equal to Matrix B's height")
	}

	var arr [][]float64

	for i := 0; i < a.Height; i++ {
		rpw := []float64{}

		for j := 0; j < b.Width; j++ {
			sum := 0.0

			for k := 0; k < a.Width; k++ {
				sum += a.Matrix[i][k] * b.Matrix[k][j]
			}

			rpw = append(rpw, sum)
		}

		arr = append(arr, rpw)
	}

	return Matrix{
		Height: len(arr),
		Width:  len(arr[0]),
		Matrix: arr,
	}, nil
}

func (a Matrix) Div(b Matrix) (Matrix, error) {
	if !canBeMultiplied(a, b) {
		return Matrix{}, errors.New("Matrix A's width must be equal to Matrix B's height")
	}

	var arr [][]float64

	for i := 0; i < a.Height; i++ {
		rpw := []float64{}

		for j := 0; j < b.Width; j++ {
			sum := 0.0

			for k := 0; k < a.Width; k++ {
				sum += a.Matrix[i][k] / b.Matrix[k][j]
			}

			rpw = append(rpw, sum)
		}

		arr = append(arr, rpw)
	}

	return Matrix{
		Height: len(arr),
		Width:  len(arr[0]),
		Matrix: arr,
	}, nil
}

func (a Matrix) Equals(b Matrix) bool {
	if a.Height != b.Height || a.Width != b.Width {
		return false
	}

	for i, j := range a.Matrix {
		for k, l := range j {
			if l != b.Matrix[i][k] {
				return false
			}
		}
	}

	return true
}

func (a Matrix) ToString() string {
	str := ""

	for _, i := range a.Matrix {
		str += fmt.Sprintf("%v\n", i)
	}

	return str
}
