package main

import (
	"fmt"
)

func main() {
	fmt.Print("Hello, playground")
	var num string
	fmt.Scanln(&num)
	
	fmt.Println(num)
}