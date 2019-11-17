package main

import (
  "log"    ///Must Include
  "os"     ///These Imports
)

func main() {
file, err := os.Open("something.txt") //Declare "err" with var
//^this "file"

  if err != nil {               //If statement checks if their is no error "err != nil"
    log.Fatal(err)              //Using the log. import, log a "Fatal Error"
  }
  defer file.Close()            //Outside the if statement defer "file.Close()"
 }     //^Goes here
