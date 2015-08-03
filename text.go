package tokeq

import "golang.org/x/net/html"

// FindText - finds text in given node
func FindText(input *html.Node) (data string) {
	for c := input.FirstChild; c != nil; c = c.NextSibling {
		if c.Parent.DataAtom == input.DataAtom && c.Type == html.TextNode {
			data += c.Data
		}
		FindText(c)
	}
	return data
}

// FindDeepText - finds text in given & child nodes
func FindDeepText(n *html.Node) (data string) {
	var f func(*html.Node)
	f = func(input *html.Node) {
		for c := input.FirstChild; c != nil; c = c.NextSibling {
			if c.Type == html.TextNode {
				data += c.Data
			}
			f(c)
		}
	}
	f(n)
	return data
}
