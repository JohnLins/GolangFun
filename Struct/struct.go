package main

import (
	"fmt"
)

type Doctor struct {
	Number int
	ActorName string
	Companions []string
}

func main() {
	aDoctor := Doctor{
		Number: 3,
		ActorName: "Jon Pertwee",
		Companions: []string {
			"Liz",
			"Jo",
			"Paul",
		},
	}
	fmt.Println(aDoctor.ActorName)

	//Anonymous Struct
	thing := struct{name string}{name: "John Lins"}
	fmt.Println(name)
}