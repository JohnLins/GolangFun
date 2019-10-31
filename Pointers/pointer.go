package main

import "fmt"


type myStruct struct {
	foo int
}

func main() {
	var a int = 42
	var b *int = &a // The "*" is Declaring pointer to an integer
	fmt.Println(a, *b, b) // The "*" is dereferencing pointer by getting value
	a = 2 //New value
	fmt.Println(a, *b, b)


	fmt.Println("____________")

	//Another Example
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

	fmt.Println("____________")

	var ms *myStruct
	ms = new(myStruct)
	ms.foo = 42
	fmt.Println(ms.foo)

	fmt.Println("____________")

	d := []int{1, 2, 3} //SLICES HAVE POINTERS BUILT IN, (Maps do also)
	f := d
	fmt.Println(d, f)
	d[1] = 42
	fmt.Println(d, f)
}


