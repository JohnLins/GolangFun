package main

import (
	"fmt"
	"encoding/json"
)

func main() {
	str := `{"name": "battery sensor", "capacity": 40, "time": "2019-01-21T19:07:28Z"}`

	var reading map[string]interface{}
	err := json.Unmarshal([]byte(str), &reading)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", reading)
}