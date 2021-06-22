package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
a := []byte(`


			<html>
<p href="23r3">                weeff we</p>          <h1>ergegr</h1>
</html> `)
//fmt.Println(GetHtml("http://paypal.com/"))
	b := CompressHTML(a)
	c := CompressHTML(b)
	fmt.Println(string(a))
	fmt.Println("____________")
	fmt.Println(string(b))
	fmt.Println("____________")
	fmt.Println(string(c))
	CreateFile(a)
}

func CompressHTML(a []byte) []byte{
	for i := 0; i < len(a) - 1; i++ {
		//Get rid of indents(10) and tabs(9)
		if a[i] == 10 || a[i] == 9{
			copy(a[i:], a[i+1:])
			a[len(a)-1] = 0
			a = a[:len(a)-1]
		}
		//Get rid of spaces(32) if the rune after it is also a space, resulting in one space.
		if a[i] == 32 && a[i + 1] == 32{
			for j := 0; ; j++ {
				if a[i + j] == 32 {
					k := i + j
				if a[k + 1] == 32 {
					k+= 1
			}
			copy(a[k:], a[k+1:])
			a[len(a)-1] = 0
			a = a[:len(a)-1]
			} else {
				break
				}
			}
		}
	}

	return a
}

func CreateFile(input []byte) {
	file, err := os.Create("TEST.html") // Truncates if file already exists, be careful!
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	defer file.Close() // Make sure to close the file when you're done

	len, err := file.WriteString(string(input))

	if err != nil {
		log.Fatalf("failed writing to file: %s", err)
	}
	fmt.Printf("\nLength: %d bytes", len)
	fmt.Printf("\nFile Name: %s", file.Name())
}


func GetHtml(url string) []byte {
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
