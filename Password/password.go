package main

import (
	"fmt"
)

func program() {
  fmt.Println("YAY")
}

func main() {
  var password string
  var username string

	fmt.Print("Create a Password: ")
  	fmt.Scanln(&password)
  fmt.Print("Create a username: ")
  	fmt.Scanln(&username)

var Apassword string
var Ausername string

  fmt.Println("Please enter your password: ")
    fmt.Scanln(&Apassword)
  fmt.Println("Please enter your username: ")
    fmt.Scanln(&Ausername)

  if Apassword == password && Ausername == username{
    program()
  } else {
    fmt.Println("YOUR WRONG!!!")
  }

}