package main

import (
	"fmt"
)

func main() {
	a := make([]int, 3, 100)
	fmt.Println(a)
	fmt.Printf("A Length: %v\n", len(a))
	fmt.Printf("A Capacity: %v\n", cap(a))
	
	/*Slice before the append*/     fmt.Println("________________")
	b := []int{}
	fmt.Println(b)
	fmt.Printf("B Length: %v\n", len(b))
	fmt.Printf("B Capacity: %v\n", cap(b))

	/*Append the slice*/            fmt.Println("________________")
	c := []int{}
	c = append(c, 1)
	fmt.Println(c)
	fmt.Printf("C Length: %v\n", len(c))
	fmt.Printf("C Capacity: %v\n", cap(c))

	/*Append but longer*/           fmt.Println("________________")
	d := []int{}
	d = append(d, 2, 4, 7, 8, 10)
	fmt.Println(d)
	fmt.Printf("D Length: %v\n", len(d))
	fmt.Printf("D Capacity: %v\n", cap(d))

	/*Link it together*/            fmt.Println("________________")
	e := []int{}
	e = append(e, []int{2, 4, 7, 8, 10}...)
	fmt.Println(e)
	fmt.Printf("E Length: %v\n", len(e))
	fmt.Printf("E Capacity: %v\n", cap(e))

	/*Take first thing off stack*/  fmt.Println("________________")
	f := []int{1, 2, 3, 4, 5}
	ff := f[1:] //Take "1" off
	fmt.Println(ff)

	/*Take last things off stack*/  fmt.Println("________________")
	g := []int{1, 2, 3, 4, 5}
	gg := g[:len(g) - 1] //Take off last
	fmt.Println(gg)

	/*Take middle things off stack*/fmt.Println("________________")
	h := []int{1, 2, 3, 4, 5}
	hh := append(h[:2], h[3:]...) //Take off "3" in middle
	fmt.Println(hh)
	
}