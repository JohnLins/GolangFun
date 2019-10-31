package main

import (
	"fmt"
)

func main() {
	statePopulations := make(map[string]int)
	//                      Key    Value            
	//                      type   Type
	statePopulations = map[string]int{
		"California":      39250017,
		"Texas":           27862596,
		"Florida":         20612439,
		"New York":        19745289,
		"Pennsylvania":    12802503,
		"Illinois":        12801539,
		"Ohio":            11614373,
	}
	statePopulations["Georgia"] = 10310371
	delete(statePopulations, "New York")
	fmt.Println(statePopulations["Texas"])
	fmt.Println(statePopulations, "\n")

	//Print map with loop and range
	for k, v := range statePopulations {
		fmt.Println(k, v)
	}
}