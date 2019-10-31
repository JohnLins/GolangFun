package main

import (
	"fmt"
)

func main() {
	var i interface{} = '0'
	switch i.(type) {
	case int:
		fmt.Println("I is an integer")
	case string:
		fmt.Println("I is a string")
	case rune:
		fmt.Println("I is a rune")
	default:
		fmt.Println("Error")
	}
}