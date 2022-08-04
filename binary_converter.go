// Binary to decimal converter code
package main

import "fmt"

//func divides decimal number to 2 and add the reminder in to digits array until decimal cannot divided
func decimalToBinary(decimalNumber int) {
	var digits []int
	for decimalNumber != 0 {
		var integer int = decimalNumber % 2
		decimalNumber = decimalNumber / 2
		digits[0] = integer

	}
	fmt.Println(digits)
}

func main() {
	var num int
	fmt.Println("Enter the number you want to convert :")
	fmt.Scan(&num)
	decimalToBinary(num)

}
