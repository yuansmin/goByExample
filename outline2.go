// Outline prints the outline of an HTML document tree.
// 打印网页的轮廓
// usage: go run outline2.go <url>
package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	//!+call
	forEachNode(doc, startElement, endElement)
	//!-call

	return nil
}

//!+forEachNode
// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, pre, post func(*html.Node, *int)) {
	var depth int
	var nodeHandle func(*html.Node)
	nodeHandle = func(n *html.Node) {
		if pre != nil {
			pre(n, &depth)
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			nodeHandle(c)
		}

		if post != nil {
			post(n, &depth)
		}
	}
	nodeHandle(n)
}

//!-forEachNode

func startElement(n *html.Node, depth *int) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", (*depth)*2, "", n.Data)
		(*depth)++
	}
}

func endElement(n *html.Node, depth *int) {
	if n.Type == html.ElementNode {
		(*depth)--
		fmt.Printf("%*s</%s>\n", (*depth)*2, "", n.Data)
	}
}
