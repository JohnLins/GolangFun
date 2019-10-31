package main

import (
	"fmt"
)

func main() {
	//Pointer with Map
	a := map[string]string {"foo": "bar", "baz": "buz"}
	b := a
	fmt.Println(a, b)
	a["foo"] = "qux"
	fmt.Println(a, b)
}