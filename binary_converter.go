// Binary to decimal converter code
package main

import "fmt"

//func divides decimal number to 2 and add the reminder in to digits array until decimal cannot divided
func decimalToBinary(decimalNumber int) {
	var digits []int
	for decimalNumber != 0 {
		var integer int = decimalNumber % 2
		digits[0] = integer
		decimalNumber = decimalNumber / 2

	}
	fmt.Println(digits)
}
func binaryToDecimal(number string) {
	// Assuming that number contains 0,1s
	// Used to store result
	var result int64 = 0
	var bit int = 0
	var n int = len(number) - 1
	// Execute given number in reverse order
	for n >= 0 {
		if number[n] == '1' {
			// When get binary 1
			result += (1 << (bit))
		}
		n = n - 1
		// Count number of bits
		bit++
	}
	// Display decimal result
	fmt.Println("  Decimal :  ", result)
}

func main() {
	var num int
	fmt.Println("Enter the number you want to convert :")
	fmt.Scan(&num)
	decimalToBinary(num)

}
