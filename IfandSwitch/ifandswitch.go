package main

import (
	"fmt"
)

func main() {

	///////////IF, ELSE IF, ELSE////////////
	var u int = 2
	if u == 1 {
		fmt.Println("1")
	} else if u == 2 {
		fmt.Println("2")
	} else {
		fmt.Println("Other")
	}

	///////////SWITCH////////////
	switch 2 {
	case 1:
		fmt.Println("1")
	case 2: 
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	default:
		fmt.Println("Other")
	}

	///////////SWITCH WITH MULTIPLE CASES////////////
	switch 9 {
	case 1, 4, 7, 9:
		fmt.Println("1")
	case 2, 6, 3, 5: 
		fmt.Println("2")
	default:
		fmt.Println("Other")
	}

	///////////TAGLESS SWITCH////////////
	var i int = 10
	switch {
	case i <= 10:
		fmt.Println("Less than or equal to ten")
	case i <= 20:
		fmt.Println("Less than or equal to twenty")
	default:
		fmt.Println("Greator than twenty")
	}
	
	///////////TYPE SWITCH////////////
	var a interface{} = 1
	switch a.(type) {
	case int:
		fmt.Println("Its an int")
	case float64:
		fmt.Println("Its a float64")
	case string:
		fmt.Println("Its a string")
	default:
		fmt.Println("Its another type")
	}

	///////////SWITCH USING FALLTHROUGH////////////
	var ii int = 10
	switch {
	case ii <= 10:
		fmt.Println("Less than or equal to ten")
		fallthrough //It will fallthrough to case 2
	case ii <= 20:
		fmt.Println("Less than or equal to twenty")
	default:
		fmt.Println("Greator than twenty")
	}

	///////////SWITCH WITH BREAK////////////
	switch 3 {
	case 1:
		fmt.Println("1")
	break
		fmt.Println("This will not print because of the break!")
	case 2: 
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	default:
		fmt.Println("Other")
	}
}