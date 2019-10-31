package main

import (
	"fmt"
)

func main() {
	myFunc("hi")
	twoArgsWithSameType("hi", "hello")

	//PASSING POINTERS
	name := "jeff"
	funcWithPointers(&name)

	//VARIADIC PARAMETER
	sum(1, 2, 3, 4, 5)

	//Func with return value
	myFuncReturn(2, 6)

	//Func returns pointer
	myFuncReturnAsPointer(6)

	//Result before end
	funcResult(7)

	//Returns two values             
	d, err := divide(5.0, 3.0)
	if err != nil {
		fmt.Println(err)
		return            //^ Basic error checking ^
	}
	fmt.Println(d)


	//Anonymous Func example
	var j int = 7
	for i := 0; i <= 5; i++ {
		func(j, i int) {
			fmt.Println(i + j)
		}(i, j)
	}

	//Anonymous Func as VAR example
	f := func(input string) {
		fmt.Println(input)
	}

	f("hi")  //Call on func

	//Method example
	g := greeter {
		greeting: "hello",
		name: "go",
	}
	g.greet()
}

func myFunc(msg string) {
	fmt.Println(msg)
}

func twoArgsWithSameType(greeting, name string) { //(greeting, name string) NOT (greeting string, name string)
	fmt.Println(greeting, name)
}

//Func with pointers
func funcWithPointers(yourName *string) {
	*yourName = "John"
	fmt.Println(*yourName)
}

//VARIADIC PARAMETER
func sum(values ...int) {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	fmt.Println("The sum is: ", result)
}

//Func return value           //Return type
func myFuncReturn(num1, num2 int) int{
	var result int = num1 + num2
	return result
}

//Func return value AS POINTER     //Return type
func myFuncReturnAsPointer(num int) *int {
	result := num * 2
	return &result
}

//Return declaired before using RESULT
func funcResult(num int) (result int) {
	result = num
	return 
}

//This func return 2 values
func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}
	return a / b, nil
}


//Method
type greeter struct {
	greeting string
	name string
}

func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
}