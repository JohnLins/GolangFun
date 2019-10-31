package main

import (
	"fmt"
)

func main() {
	a := 10
	b := 3

	fmt.Println( a + b)
	fmt.Println( a - b)
	fmt.Println( a * b)
	fmt.Println( a / b)
	fmt.Println( a % b)

	fmt.Println(a & b) // 0010
	fmt.Println(a | b) // 1011
	fmt.Println(a ^ b) // 1001
	fmt.Println(a &^ b)// 0100

	//Bit shifting
	c := 8 //2^3
	fmt.Println(c << 3) // 2^3 * 2^3 = 2^6 = 64
	fmt.Println(c >> 3) // 2^3 / 2^3 = 2^0 = 1
}