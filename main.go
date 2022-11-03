package main

import (
	"fmt"
	"strconv"
)

func main() {
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
}
