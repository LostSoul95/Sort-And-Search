package main

import (
	"Binary/pkg"
	"fmt"
)

func main() {

	var inputSize int
	fmt.Println("Enter the size of the array you want ")
	fmt.Scan(&inputSize)

	var inputNumber int
	fmt.Println("Enter number you want to find :")
	fmt.Scan(&inputNumber)

	pkg.ExecuteBinarySearch(inputSize, inputNumber)
	//fmt.Println("read number", i, "from stdin")

}
