package crawler

import (
    "bytes"
    "errors"
    "fmt"
    "golang.org/x/net/html"
    "io"
    "strings"
    "io/ioutil"
    "net/http"
)
/*
func main() {
    URLCrawl("script", "https://www.google.com/")
}
*/
func Element(doc *html.Node, tagName string) (*html.Node, error) {
    var element *html.Node
    var crawler func(*html.Node)
    crawler = func(node *html.Node) {
        if node.Type == html.ElementNode && node.Data == tagName {
            element = node
            return
        }
        for child := node.FirstChild; child != nil; child = child.NextSibling {
            crawler(child)
        }
    }
    crawler(doc)
    if element != nil {
        return element, nil
    }
    return nil, errors.New("Missing in the node tree")
}

func renderNode(n *html.Node) string {
    var buf bytes.Buffer
    w := io.Writer(&buf)
    html.Render(w, n)
    return buf.String()
}

func URLCrawl(findElement, url string)[]byte {
    resp, err := http.Get(url)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
    sitehtml, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }

    doc, _ := html.Parse(strings.NewReader(string(sitehtml)))
    bn, err := Element(doc, findElement)
    //if err != nil {
        //return
    //}
    element := renderNode(bn)
    fmt.Println(element)
    return []byte(element)
}
