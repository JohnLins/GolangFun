package main

import (
	"fmt"
)

func main() {
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello Go"))
}

type Writer interface {
	Write([]byte) (int, error)
				  //Return
}

type ConsoleWriter struct {}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

//ADVICE
//Make interfaces as small as possible to maximize scalability