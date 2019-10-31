package main

import (
	"fmt"
)

const (
	_ = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	var fileSize float64 = 4000000000
	fmt.Println(fileSize/KB, "KB") 
	fmt.Println(fileSize/MB, "MB")
	fmt.Println(fileSize/GB, "GB")
	fmt.Println(fileSize/TB, "TB")
	fmt.Println(fileSize/PB, "PB")
	fmt.Println(fileSize/EB, "EB")
	fmt.Println(fileSize/ZB, "ZB")
	fmt.Println(fileSize/YB, "YB")
}