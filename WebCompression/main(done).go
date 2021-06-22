package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
)

func compressHTML(url string) ([]byte, int) {

	var ignoreJavascript bool = false

	var reduction int = 0; 
	remove := func(s []byte, i int) []byte {
		return append(s[:i], s[i+1:]...)
	}

	html := getHTML(url)
	for i := 0; i<len(html)-2; i++ {

		if string(html[i:i+7]) == "<script" {
			ignoreJavascript = true;
		}

		if string(html[i:i+8]) == "</script>" {
			ignoreJavascript = false;
		}
		
		if ignoreJavascript == false {
			//return
			if html[i] == 10 {
			html = remove(html, i)
					reduction++;
			}

			//Space
			for html[i + 1] == 32 && html[i] == 32 {
				html = remove(html, i)
					reduction++;
			}
		}
    

	}

	fmt.Println(string(html))
	fmt.Println(reduction)

		
	return html, reduction
}

func main() {
var compressedHTML, reduction = compressHTML("http://linsintegrations.co")	
fmt.Println(string(compressedHTML))
fmt.Println(reduction, "B")	 
}

func getHTML(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return html
}
