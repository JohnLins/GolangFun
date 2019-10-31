package main

import (
	"fmt"
)

type Animal struct {
	Name string
	Origin string
}

type Bird struct {
	Animal //Bird is inherating from Animal
	SpeedKPH  float32
	CanFly bool
}

func main() {
	b := Bird{
		Animal: Animal{
			Name:    "Emu",
			Origin:  "Australia",
		},
		SpeedKPH:     48,
		CanFly:       false,
	}
	fmt.Println(b.Name)
}