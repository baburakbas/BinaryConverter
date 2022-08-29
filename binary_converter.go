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

func main() {
	// we take decimal number from user and use it in decimalToBinary function 
	var numDeci int
	fmt.Print("Enter the number you want to convert :")
	fmt.Scan(&numDeci)
	decimalToBinary(numDeci)
	hello()

}
func hello(){
	fmt.Println("Merhaba Hasan")

}

