package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string `required max: "100"` //This is a tag, this tag sets the max to 100
	Age int
}

func main() {
	t := reflect.TypeOf(User{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}