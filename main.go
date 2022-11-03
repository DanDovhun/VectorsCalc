package main

import (
	"fmt"
	"strconv"
)

func main() {
	options := []string{"2D Vector", "3D Vector", "Matrix"}

	getInt := func(msg string) int64 {
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

	getFloat := func(msg string) float64 {
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

	for {
		fmt.Println("Options: ")

		for i, j := range options {
			fmt.Printf("%v.) %v\n", i+1, j)
		}

		opt := getInt("Choice: ")

		switch opt {
		case 1:
			fmt.Println("2D Vectors")

		case 2:
			fmt.Println("3D Vectors")

		case 3:
			fmt.Println("Matrix")

		default:
			fmt.Println("Invalid option")
		}
	}
}
