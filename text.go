package tokeq

import (
	"strings"

	"golang.org/x/net/html"
)

// FindText - finds text in given node
func FindText(input *html.Node) (data string) {
	var buffer string
	for c := input.FirstChild; c != nil; c = c.NextSibling {
		if c.Parent.DataAtom == input.DataAtom && c.Type == html.TextNode {
			buffer = strings.TrimSpace(c.Data)
			if len(buffer) > 0 {
				data += " " + buffer
			}
		}
		FindText(c)
	}
	return strings.TrimSpace(data)
}

// FindDeepText - finds text in given & child nodes
func FindDeepText(n *html.Node) (data string) {
	var buffer string
	var f func(*html.Node)
	f = func(input *html.Node) {
		for c := input.FirstChild; c != nil; c = c.NextSibling {
			if c.Type == html.TextNode {
				buffer = strings.TrimSpace(c.Data)
				if len(buffer) > 0 {
					data += " " + buffer
				}
			}
			f(c)
		}
	}
	f(n)
	return strings.TrimSpace(data)
}
