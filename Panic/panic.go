package main

import (
	"fmt"
)

func main() {
	fmt.Println("Start")
	panic("Something bad happened")
	fmt.Println("end")
}