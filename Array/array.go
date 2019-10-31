package main

import (
	"fmt"
)

func main() {
	//Arrays & slices

	//Static Array
	numsstatic := [3]int{97, 85, 93}

	//Unspecified
	var wordsunspecified [3]string
	wordsunspecified[1] = "John"

	var mulitiseg [3][3]int = [3][3]int{ 
	[3]int{1, 0, 0}, 
	[3]int{0, 1, 0},
	[3]int{0, 0, 1}}

	//Slice
		a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		b := a[:]
		c := a[3:]
		d := a[:6]
		e := a[3:6]

		fmt.Println(a, "   ", b, "   ", c, "   ", d, "   ", e)



	fmt.Println(numsstatic[1])
	fmt.Println(wordsunspecified[1])
	fmt.Println(mulitiseg)

	for i := 0; i <= 2; i++ {
		fmt.Println(numsstatic[i])
	}


	//A DYNAMIC ARRAY:

	arr := [...]int{1, 2, 3}
	fmt.Println(arr)
}

