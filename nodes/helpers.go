package nodes

import (
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// ParentIsHead - whether parent is not nil and has <head> tag as parent node
func ParentIsHead(n *html.Node) bool {
	return n != nil && n.Parent != nil && n.Parent.DataAtom == atom.Head
}

// ParentIsBody - whether parent is not nil and has <body> tag as parent node
func ParentIsBody(n *html.Node) bool {
	return n != nil && n.Parent != nil && n.Parent.DataAtom == atom.Body
}
