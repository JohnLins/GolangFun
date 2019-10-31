package main

import (
    "fmt"
    "math/rand"
    "time"
)

func random(min int, max int) int {
    rand.Seed(time.Now().UnixNano())
    return rand.Intn(max-min) + min
}

func main() {              //MIN  MAX
        randomNum := random(1000, 2000)
        fmt.Println(randomNum)
}