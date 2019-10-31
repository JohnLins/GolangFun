package main

import (
	"fmt"
)

func main() {
	//Basic For Loop
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}

	//Loop with more than one values
	for a, j := 0, 0; a <= 5; a, j = a+1, j+2 {
		fmt.Println(a, j)
	}

	//While Loop
	var b int = 3
	for ; b <= 9; b++ {
		fmt.Println(b)
	}

	//While with Incrementor inside
	var c int = 7
	for ; c <= 19 ; {
		c++
		fmt.Println(c)
	}

	//Loop with break
	for d := 0; d <= 5; d++ {
		fmt.Println(d)
		if d == 4 {
			break
		}
	}

	//Loop with continue
	for e := 0; e <= 10; e++ {
		if e%2 == 0 {
			continue
		}
		fmt.Println(e)
	}

	//Nested Loop
	for f := 1; f <= 3; f++ {
		for g := 1; g <= 3; g++ {
			fmt.Println(f * g)
		}
	}

	//Nested Loop with label
loop: 
	for p := 1; p <= 3; p++ {
		for q := 1; q <= 3; q++ {
			fmt.Println(p * q)
			if p * q >= 3 {
				break loop
			}
		}
	}

	//Loop with range
	s := []int{1, 2, 3}
	for k, v := range s {
		fmt.Println("Location:", k, "value:", v)
	}

	//Loop without key using Go's universal placeholder '_'
	x := []int{1, 2, 3}
	for _, y := range x {
		fmt.Println("value:", y)
	}

	//Another example with range
	h := "Hello"
	for m, n := range h {
		fmt.Println(m, n) //ASCII
		fmt.Println(m, string(n)) //Text
	}
}