package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/DanDovhun/VectorsCalc/three_dim"
	"github.com/DanDovhun/VectorsCalc/two_dim"
)

var matrix = func() {
	fmt.Println("Options:")
	var options = []string{"Add", "Subtract", "Multiply", "Divide"}

	for i, j := range options {
		fmt.Printf("%v) %v\n", i+1, j)
	}

	fmt.Print("Choice")
}

var getInt = func(msg string) int64 {
	for {
		var str string

		fmt.Print(msg)
		fmt.Scanln(&str)

		num, err := strconv.ParseInt(str, 10, 64)

		if err != nil {
			fmt.Println("Please only enter whole numbers")
			continue
		}

		return num
	}
}

var getFloat = func(msg string) float64 {
	for {
		var str string

		fmt.Print(msg)
		fmt.Scanln(&str)

		num, err := strconv.ParseFloat(str, 64)

		if err != nil {
			fmt.Println("Please only enter whole numbers")
			continue
		}

		return num
	}
}

func main() {
	options := []string{"2D Vector", "3D Vector", "Matrix"}

	round := func(num float64, to int) float64 {
		return math.Round(num*math.Pow10(to)) / math.Pow10(to)
	}

	for {
		fmt.Println("Options: ")

		for i, j := range options {
			fmt.Printf("%v.) %v\n", i+1, j)
		}

		opt := getInt("Choice: ")

		switch opt {
		case 1:
			fmt.Println("\n2D Vector")

			fmt.Println("Vector A:")
			xA := getFloat("x: ")
			yA := getFloat("y: ")

			fmt.Println("\nVector B:")
			xB := getFloat("x: ")
			yB := getFloat("y: ")

			vecA := two_dim.New("A", xA, yA)
			vecB := two_dim.New("B", xB, yB)

			fmt.Println("\nVector A: " + vecA.ToString())
			fmt.Printf("Magnitude: %v\n", round(vecA.Magnitude(), 5))
			fmt.Printf("Elevation: %v radians; %v degrees\n", round(vecA.Elevation().Radians, 5), round(vecA.Elevation().Degrees, 5))

			fmt.Println("\nVector B: " + vecB.ToString())
			fmt.Printf("Magnitude: %v\n", round(vecB.Magnitude(), 5))
			fmt.Printf("Elevation: %v radians; %v degrees\n", round(vecB.Elevation().Radians, 5), round(vecB.Elevation().Degrees, 5))

			fmt.Println("\nRelationships:")

			add := vecA.Add(vecB)
			sub := vecA.Sub(vecB)
			dProd := vecA.DotProduct(vecB)
			cProd := vecA.CrossProduct(vecB)
			angleBetween := vecA.AngleBetween(vecB)

			fmt.Printf("A + B: %v\n", add.ToString())
			fmt.Printf("A - B: %v\n", sub.ToString())
			fmt.Printf("Dot Product: %v\n", dProd)
			fmt.Printf("Cross Product: %v\n", cProd)
			fmt.Printf("Angle between A and B: %v radians, %v degrees\n", round(angleBetween.Radians, 5), round(angleBetween.Degrees, 5))

		case 2:
			fmt.Println("\n3D Vectors")

			fmt.Println("Vector A:")
			xA := getFloat("X: ")
			yA := getFloat("Y: ")
			zA := getFloat("Z: ")

			fmt.Println("\nVector B:")
			xB := getFloat("X: ")
			yB := getFloat("Y: ")
			zB := getFloat("Z: ")

			vecA := three_dim.New("A", xA, yA, zA)
			vecB := three_dim.New("A", xB, yB, zB)

			fmt.Printf("\nVector A: %v\n", vecA.ToString())
			paramsA := vecA.GetProperties()
			alphaA, betaA, gammaA, elevationA, rotationA := paramsA.GetProperties()

			fmt.Printf("Magnitude: %v\n", round(vecA.Magnitude(), 5))
			fmt.Printf("Alpha: %v radians; %v degrees\n", round(alphaA.Radians, 5), round(alphaA.Degrees, 5))
			fmt.Printf("Beta: %v radians; %v degrees\n", round(betaA.Radians, 5), round(betaA.Degrees, 5))
			fmt.Printf("Gamma: %v radians; %v degrees\n", round(gammaA.Radians, 5), round(gammaA.Degrees, 5))
			fmt.Printf("Elevation: %v radians; %v degrees\n", round(elevationA.Radians, 5), round(elevationA.Degrees, 5))
			fmt.Printf("Rotation: %v radians; %v degrees\n", round(rotationA.Radians, 5), round(rotationA.Degrees, 5))

			fmt.Printf("\nVector B:%v\n", vecB.ToString())
			paramsB := vecB.GetProperties()
			alphaB, betaB, gammaB, elevationB, rotationB := paramsB.GetProperties()

			fmt.Printf("Magnitude: %v\n", round(vecB.Magnitude(), 5))
			fmt.Printf("Alpha: %v radians; %v degrees\n", round(alphaB.Radians, 5), round(alphaB.Degrees, 5))
			fmt.Printf("Beta: %v radians; %v degrees\n", round(betaB.Radians, 5), round(betaB.Degrees, 5))
			fmt.Printf("Gamma: %v radians; %v degrees\n", round(gammaB.Radians, 5), round(gammaB.Degrees, 5))
			fmt.Printf("Elevation: %v radians; %v degrees\n", round(elevationB.Radians, 5), round(elevationB.Degrees, 5))
			fmt.Printf("Rotation: %v radians; %v degrees\n", round(rotationB.Radians, 5), round(rotationB.Degrees, 5))

			fmt.Println("\nRelationships:")

			add := vecA.Add(vecB)
			sub := vecA.Sub(vecB)
			dProd := vecA.DotProduct(vecB)
			cProd := vecA.CrossProduct(vecB)
			angleBetween := vecA.AngleBetween(vecB)

			fmt.Printf("A + B: %v\n", add.ToString())
			fmt.Printf("A - B: %v\n", sub.ToString())
			fmt.Printf("Dot Product: %v\n", dProd)
			fmt.Printf("Cross Product: %v\n", cProd.ToString())
			fmt.Printf("Angle between A and B: %v radians, %v degrees\n", round(angleBetween.Radians, 5), round(angleBetween.Degrees, 5))

		case 3:
			fmt.Println("Matrix")

		default:
			fmt.Println("Invalid option")
		}

		fmt.Println("")
	}
}
