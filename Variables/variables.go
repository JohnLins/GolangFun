package main

import "fmt"
//Var block

var (
	fname string = "John"
	age int = 15
	alive bool = true
	gpa float32 = 4.2
	bit64float float64 = 1.9
	unsigned8bitint uint8 = 255
	unsigned16bitint uint16 = 65535
	unsigned32bitint uint32 = 294967295
	complexnum complex64 = 1 + 2i
	letter rune = 'a'
)
//Const block
const (
	stuff= "stuff"
)

func main() {
	//You can use := as "AUTO" in c++
	d := "text"
	const myCost int = 42

	fmt.Println(fname)
	fmt.Println(age)
	fmt.Println(alive)
	fmt.Println(gpa)
	fmt.Println(bit64float)

	fmt.Println(d)

	fmt.Println(unsigned8bitint)
	fmt.Println(unsigned16bitint)
	fmt.Println(unsigned32bitint)

	fmt.Printf("%V, %T\n", real(complexnum), real(complexnum))
	fmt.Printf("%V, %T\n", imag(complexnum), imag(complexnum))

	fmt.Println(letter)
	fmt.Println(myCost)

	fmt.Println(stuff)
}

//ALL TYPES
/*
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
*/
