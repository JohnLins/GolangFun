package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {
	resp, err := http.Get("https://www.udemy.com/")   
	
	if err != nil {
		panic(err)
	}
	
	defer resp.Body.Close()
	
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	
	fmt.Printf("%s", html)
}
