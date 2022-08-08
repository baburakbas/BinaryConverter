// Binary to decimal converter code
package main

import "fmt"

//func divides decimal number to 2 and add the reminder in to digits array until decimal cannot divided
// array has 64 digits, so we can calculate 64 digits of binary
func decimalToBinary(decimalNumber int) {
	var digits [64]int
	var digitNo int = 63
	for decimalNumber != 0 {
		var integer int = decimalNumber % 2
		decimalNumber = decimalNumber / 2
		digits[digitNo] = integer
		digitNo--

	}
	// this section limits the range we want to print
	fmt.Print("Binary equivalent of your number : ")
	for printRange := 63; printRange != digitNo; printRange-- {
		fmt.Print(digits[printRange])
	}

}

func binaryToDecimal(binaryNumber int) {
	var deciArray [64]int
	var elementNo int = 0
	for i := 0; binaryNumber != 0; i++ {
		var reminde int = binaryNumber % 10
		for n := 0; n != i; n++ {
			var elementArray int = reminde * 2
			deciArray[elementNo] = elementArray
		}

		binaryNumber = binaryNumber / 10
		elementNo++
	}
	fmt.Print(deciArray)
}

func main() {
	var numDeci int

	fmt.Println("Enter the number you want to convert :")
	fmt.Scan(&numDeci)
	decimalToBinary(numDeci)
}
