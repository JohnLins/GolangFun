package main

import (
	"fmt"
)

func main() {
	fmt.Println("Start")
	defer fmt.Println("Middle") //A defer function will run once the main func is complete!
	fmt.Println("End")


	//A DEFER FUNCTION
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error", err)
		}
	}()
	
}
//DEFER PLAYS HERE


