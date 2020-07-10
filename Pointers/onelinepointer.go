package main

import "fmt"

func main() {
	i := 42
	*&i = 21         
	fmt.Println(i)     
}
